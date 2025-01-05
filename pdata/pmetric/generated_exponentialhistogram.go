// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

// ExponentialHistogram represents the type of a metric that is calculated by aggregating
// as a ExponentialHistogram of all reported double measurements over a time interval.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExponentialHistogram function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ExponentialHistogram struct {
	orig  *otlpmetrics.ExponentialHistogram
	state *internal.State
}

func newExponentialHistogram(orig *otlpmetrics.ExponentialHistogram, state *internal.State) ExponentialHistogram {
	return ExponentialHistogram{orig: orig, state: state}
}

// NewExponentialHistogram creates a new empty ExponentialHistogram.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExponentialHistogram() ExponentialHistogram {
	state := internal.StateMutable
	return newExponentialHistogram(&otlpmetrics.ExponentialHistogram{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ExponentialHistogram) MoveTo(dest ExponentialHistogram) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.ExponentialHistogram{}
}

func (ms ExponentialHistogram) Size() int {
	return ms.orig.Size()
}

// AggregationTemporality returns the aggregationtemporality associated with this ExponentialHistogram.
func (ms ExponentialHistogram) AggregationTemporality() AggregationTemporality {
	return AggregationTemporality(ms.orig.AggregationTemporality)
}

// SetAggregationTemporality replaces the aggregationtemporality associated with this ExponentialHistogram.
func (ms ExponentialHistogram) SetAggregationTemporality(v AggregationTemporality) {
	ms.state.AssertMutable()
	ms.orig.AggregationTemporality = otlpmetrics.AggregationTemporality(v)
}

// DataPoints returns the DataPoints associated with this ExponentialHistogram.
func (ms ExponentialHistogram) DataPoints() ExponentialHistogramDataPointSlice {
	return newExponentialHistogramDataPointSlice(&ms.orig.DataPoints, ms.state)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ExponentialHistogram) CopyTo(dest ExponentialHistogram) {
	dest.state.AssertMutable()
	dest.SetAggregationTemporality(ms.AggregationTemporality())
	ms.DataPoints().CopyTo(dest.DataPoints())
}
