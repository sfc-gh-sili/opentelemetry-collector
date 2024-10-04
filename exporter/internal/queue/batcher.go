// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package queue // import "go.opentelemetry.io/collector/exporter/internal/queue"

import (
	"context"
	"math"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exporterbatcher"
	"go.opentelemetry.io/collector/exporter/internal"
)

type batch struct {
	ctx              context.Context
	req              internal.Request
	onExportFinished func(error)
}

type Batcher struct {
	cfg            exporterbatcher.Config
	mergeFunc      exporterbatcher.BatchMergeFunc[internal.Request]
	mergeSplitFunc exporterbatcher.BatchMergeSplitFunc[internal.Request]

	queue      Queue[internal.Request]
	numWorkers int
	exportFunc func(context.Context, internal.Request) error
	stopWG     sync.WaitGroup

	mu             sync.Mutex
	lastFlushed    time.Time
	pendingBatches []batch
	timer          *time.Timer
}

func NewBatcher(cfg exporterbatcher.Config, queue Queue[internal.Request], numWorkers int, exportFunc func(context.Context, internal.Request) error) *Batcher {
	return &Batcher{
		cfg:            cfg,
		queue:          queue,
		numWorkers:     numWorkers,
		exportFunc:     exportFunc,
		stopWG:         sync.WaitGroup{},
		pendingBatches: make([]batch, 1),
	}
}

func (qb *Batcher) flushIfNecessary() {
	qb.mu.Lock()

	if qb.pendingBatches[0].req == nil {
		qb.mu.Unlock()
		return
	}

	if time.Since(qb.lastFlushed) < qb.cfg.FlushTimeout && qb.pendingBatches[0].req.ItemsCount() < qb.cfg.MinSizeItems {
		qb.mu.Unlock()
		return
	}

	flushedBatch := qb.pendingBatches[0]
	qb.pendingBatches = qb.pendingBatches[1:]
	if len(qb.pendingBatches) == 0 {
		qb.pendingBatches = append(qb.pendingBatches, batch{})
	}

	qb.lastFlushed = time.Now()

	if qb.cfg.FlushTimeout > 0 {
		qb.timer.Reset(qb.cfg.FlushTimeout)
	}

	qb.mu.Unlock()

	err := qb.exportFunc(flushedBatch.ctx, flushedBatch.req)
	if flushedBatch.onExportFinished != nil {
		flushedBatch.onExportFinished(err)
	}
}

func (qb *Batcher) push(req internal.Request, onExportFinished func(error)) error {
	qb.mu.Lock()
	defer qb.mu.Unlock()

	idx := len(qb.pendingBatches) - 1
	if qb.pendingBatches[idx].req == nil {
		qb.pendingBatches[idx].req = req
		qb.pendingBatches[idx].ctx = context.Background()
		qb.pendingBatches[idx].onExportFinished = onExportFinished
	} else {
		reqs, err := qb.mergeSplitFunc(context.Background(),
			qb.cfg.MaxSizeConfig,
			qb.pendingBatches[idx].req, req)
		if err != nil || len(reqs) == 0 {
			return err
		}

		for offset, newReq := range reqs {
			if offset != 0 {
				qb.pendingBatches = append(qb.pendingBatches, batch{})
			}
			qb.pendingBatches[idx+offset].req = newReq
		}
	}
	return nil
}

// Start ensures that queue and all consumers are started.
func (qb *Batcher) Start(ctx context.Context, host component.Host) error {
	if err := qb.queue.Start(ctx, host); err != nil {
		return err
	}

	if qb.cfg.FlushTimeout > 0 {
		qb.timer = time.NewTimer(qb.cfg.FlushTimeout)
	} else {
		qb.timer = time.NewTimer(math.MaxInt)
		qb.timer.Stop()
	}

	allocReader := make(chan bool, 10)

	go func() {
		allocReader <- true
	}()

	var startWG sync.WaitGroup
	for i := 0; i < qb.numWorkers; i++ {
		qb.stopWG.Add(1)
		startWG.Add(1)
		go func() {
			startWG.Done()
			defer qb.stopWG.Done()
			for {
				select {
				case <-qb.timer.C:
					qb.flushIfNecessary()
				case <-allocReader:
					req, ok, onProcessingFinished := qb.queue.ClaimAndRead(func() {
						allocReader <- true
					})
					if !ok {
						return
					}

					qb.push(req, onProcessingFinished) // Handle error
					qb.flushIfNecessary()
				}
			}
		}()
	}
	startWG.Wait()

	return nil
}

// Shutdown ensures that queue and all Batcher are stopped.
func (qb *Batcher) Shutdown(ctx context.Context) error {
	if err := qb.queue.Shutdown(ctx); err != nil {
		return err
	}
	qb.stopWG.Wait()
	return nil
}
