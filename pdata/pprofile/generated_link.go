// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// Link represents a pointer from a profile Sample to a trace Span.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Link struct {
	orig  *otlpprofiles.Link
	state *internal.State
}

func newLink(orig *otlpprofiles.Link, state *internal.State) Link {
	return Link{orig: orig, state: state}
}

// NewLink creates a new empty Link.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewLink() Link {
	state := internal.StateMutable
	return newLink(&otlpprofiles.Link{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Link) MoveTo(dest Link) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlpprofiles.Link{}
}

func (ms Link) Size() int {
	return ms.orig.Size()
}

// TraceID returns the traceid associated with this Link.
func (ms Link) TraceID() pcommon.TraceID {
	return pcommon.TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this Link.
func (ms Link) SetTraceID(v pcommon.TraceID) {
	ms.state.AssertMutable()
	ms.orig.TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this Link.
func (ms Link) SpanID() pcommon.SpanID {
	return pcommon.SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this Link.
func (ms Link) SetSpanID(v pcommon.SpanID) {
	ms.state.AssertMutable()
	ms.orig.SpanId = data.SpanID(v)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Link) CopyTo(dest Link) {
	dest.state.AssertMutable()
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
}
