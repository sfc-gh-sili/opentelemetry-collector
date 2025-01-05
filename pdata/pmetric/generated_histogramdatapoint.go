// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// HistogramDataPoint is a single data point in a timeseries that describes the time-varying values of a Histogram of values.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewHistogramDataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type HistogramDataPoint struct {
	orig  *otlpmetrics.HistogramDataPoint
	state *internal.State
}

func newHistogramDataPoint(orig *otlpmetrics.HistogramDataPoint, state *internal.State) HistogramDataPoint {
	return HistogramDataPoint{orig: orig, state: state}
}

// NewHistogramDataPoint creates a new empty HistogramDataPoint.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewHistogramDataPoint() HistogramDataPoint {
	state := internal.StateMutable
	return newHistogramDataPoint(&otlpmetrics.HistogramDataPoint{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms HistogramDataPoint) MoveTo(dest HistogramDataPoint) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.HistogramDataPoint{}
}

func (ms HistogramDataPoint) Size() int {
	return ms.orig.Size()
}

// Attributes returns the Attributes associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Attributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.orig.Attributes, ms.state))
}

// StartTimestamp returns the starttimestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) StartTimestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.StartTimeUnixNano)
}

// SetStartTimestamp replaces the starttimestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetStartTimestamp(v pcommon.Timestamp) {
	ms.state.AssertMutable()
	ms.orig.StartTimeUnixNano = uint64(v)
}

// Timestamp returns the timestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Timestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetTimestamp(v pcommon.Timestamp) {
	ms.state.AssertMutable()
	ms.orig.TimeUnixNano = uint64(v)
}

// Count returns the count associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Count() uint64 {
	return ms.orig.Count
}

// SetCount replaces the count associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetCount(v uint64) {
	ms.state.AssertMutable()
	ms.orig.Count = v
}

// BucketCounts returns the bucketcounts associated with this HistogramDataPoint.
func (ms HistogramDataPoint) BucketCounts() pcommon.UInt64Slice {
	return pcommon.UInt64Slice(internal.NewUInt64Slice(&ms.orig.BucketCounts, ms.state))
}

// ExplicitBounds returns the explicitbounds associated with this HistogramDataPoint.
func (ms HistogramDataPoint) ExplicitBounds() pcommon.Float64Slice {
	return pcommon.Float64Slice(internal.NewFloat64Slice(&ms.orig.ExplicitBounds, ms.state))
}

// Exemplars returns the Exemplars associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Exemplars() ExemplarSlice {
	return newExemplarSlice(&ms.orig.Exemplars, ms.state)
}

// Flags returns the flags associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Flags() DataPointFlags {
	return DataPointFlags(ms.orig.Flags)
}

// SetFlags replaces the flags associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetFlags(v DataPointFlags) {
	ms.state.AssertMutable()
	ms.orig.Flags = uint32(v)
}

// Sum returns the sum associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Sum() float64 {
	return ms.orig.GetSum()
}

// HasSum returns true if the HistogramDataPoint contains a
// Sum value, false otherwise.
func (ms HistogramDataPoint) HasSum() bool {
	return ms.orig.Sum_ != nil
}

// SetSum replaces the sum associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetSum(v float64) {
	ms.state.AssertMutable()
	ms.orig.Sum_ = &otlpmetrics.HistogramDataPoint_Sum{Sum: v}
}

// RemoveSum removes the sum associated with this HistogramDataPoint.
func (ms HistogramDataPoint) RemoveSum() {
	ms.state.AssertMutable()
	ms.orig.Sum_ = nil
}

// Min returns the min associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Min() float64 {
	return ms.orig.GetMin()
}

// HasMin returns true if the HistogramDataPoint contains a
// Min value, false otherwise.
func (ms HistogramDataPoint) HasMin() bool {
	return ms.orig.Min_ != nil
}

// SetMin replaces the min associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetMin(v float64) {
	ms.state.AssertMutable()
	ms.orig.Min_ = &otlpmetrics.HistogramDataPoint_Min{Min: v}
}

// RemoveMin removes the min associated with this HistogramDataPoint.
func (ms HistogramDataPoint) RemoveMin() {
	ms.state.AssertMutable()
	ms.orig.Min_ = nil
}

// Max returns the max associated with this HistogramDataPoint.
func (ms HistogramDataPoint) Max() float64 {
	return ms.orig.GetMax()
}

// HasMax returns true if the HistogramDataPoint contains a
// Max value, false otherwise.
func (ms HistogramDataPoint) HasMax() bool {
	return ms.orig.Max_ != nil
}

// SetMax replaces the max associated with this HistogramDataPoint.
func (ms HistogramDataPoint) SetMax(v float64) {
	ms.state.AssertMutable()
	ms.orig.Max_ = &otlpmetrics.HistogramDataPoint_Max{Max: v}
}

// RemoveMax removes the max associated with this HistogramDataPoint.
func (ms HistogramDataPoint) RemoveMax() {
	ms.state.AssertMutable()
	ms.orig.Max_ = nil
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms HistogramDataPoint) CopyTo(dest HistogramDataPoint) {
	dest.state.AssertMutable()
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetStartTimestamp(ms.StartTimestamp())
	dest.SetTimestamp(ms.Timestamp())
	dest.SetCount(ms.Count())
	ms.BucketCounts().CopyTo(dest.BucketCounts())
	ms.ExplicitBounds().CopyTo(dest.ExplicitBounds())
	ms.Exemplars().CopyTo(dest.Exemplars())
	dest.SetFlags(ms.Flags())
	if ms.HasSum() {
		dest.SetSum(ms.Sum())
	}

	if ms.HasMin() {
		dest.SetMin(ms.Min())
	}

	if ms.HasMax() {
		dest.SetMax(ms.Max())
	}

}
