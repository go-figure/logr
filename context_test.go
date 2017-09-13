package logr_test

import (
	"context"
	"testing"

	"github.com/go-figure/logr"
	"github.com/stretchr/testify/require"
)

func TestSinkContext(t *testing.T) {
	ctx := context.Background()

	sink := logr.SinkFromContext(ctx)
	require.Equal(t, logr.Discard, sink, "SinkFromContext on an empty context should return Discard sink")

	someSink := &logr.DiscardSink{}
	ctx = logr.SinkContext(ctx, someSink)

	sink = logr.SinkFromContext(ctx)
	require.NotEqual(t, logr.Discard, sink, "SinkFromContext on an initialized context should not return Discard")
	require.Equal(t, someSink, sink, "SinkFromContext on an initialized context should return the set sink")

	require.Panics(t, func() {
		logr.SinkContext(ctx, nil)
	}, "SinkContext should panic with nil sink")
}

func TestReceiverContext(t *testing.T) {
	ctx := context.Background()

	receiver := logr.FromContext(ctx)
	require.Equal(t, logr.DiscardReceiver, receiver, "FromContext on an empty context should return DiscardReceiver")

	someReceiver := logr.NewJob(nil, "testing", nil)
	ctx = logr.ReceiverContext(ctx, someReceiver)

	receiver = logr.FromContext(ctx)
	require.NotEqual(t, logr.Discard, receiver, "FromContext on an initialized context should not return Discard receiver")
	require.Equal(t, someReceiver, receiver, "FromContext on an initialized context should return the set receiver")

	require.Panics(t, func() {
		logr.ReceiverContext(ctx, nil)
	}, "ReceiverContext should panic with nil sink")
}
