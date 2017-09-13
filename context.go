package logr

import "context"

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "logr " + k.name }

var (
	ctxSinkKey     = &contextKey{"sink"}
	ctxReceiverKey = &contextKey{"receiver"}
)

func SinkContext(ctx context.Context, sink Sink) context.Context {
	if sink == nil {
		panic("sink is nil")
	}

	return context.WithValue(ctx, ctxSinkKey, sink)
}

func SinkFromContext(ctx context.Context) Sink {
	sink, _ := ctx.Value(ctxSinkKey).(Sink)
	if sink == nil {
		return Discard
	}

	return sink
}

func ReceiverContext(ctx context.Context, rcv Receiver) context.Context {
	if rcv == nil {
		panic("receiver is nil")
	}

	return context.WithValue(ctx, ctxReceiverKey, rcv)
}

func FromContext(ctx context.Context) Receiver {
	rcv, _ := ctx.Value(ctxReceiverKey).(Receiver)
	if rcv == nil {
		return DiscardReceiver
	}

	return rcv
}
