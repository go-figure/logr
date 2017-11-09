package logr

import "time"

// check that all types implement the Sink interface
var _ Sink = DiscardSink{}
var _ Sink = TeeSink{}
var _ Sink = (*KVSink)(nil)

var Discard = DiscardSink{}

type DiscardSink struct{}

func (_ DiscardSink) Event(job, event string, kv KV)                                            {}
func (_ DiscardSink) Error(job, event string, err error, kv KV)                                 {}
func (_ DiscardSink) Timing(job, event string, timing time.Duration, kv KV)                     {}
func (_ DiscardSink) Gauge(job, event string, gauge float64, kv KV)                             {}
func (_ DiscardSink) Complete(job string, status CompletionStatus, timing time.Duration, kv KV) {}

type TeeSink []Sink

func (ts *TeeSink) Add(sink Sink) {
	*ts = append(*ts, sink)
}

func (ts TeeSink) Event(job, event string, kv KV) {
	for _, sink := range ts {
		sink.Event(job, event, kv)
	}
}

func (ts TeeSink) Error(job, event string, err error, kv KV) {
	for _, sink := range ts {
		sink.Error(job, event, err, kv)
	}
}

func (ts TeeSink) Timing(job, event string, timing time.Duration, kv KV) {
	for _, sink := range ts {
		sink.Timing(job, event, timing, kv)
	}
}

func (ts TeeSink) Gauge(job, event string, gauge float64, kv KV) {
	for _, sink := range ts {
		sink.Gauge(job, event, gauge, kv)
	}
}

func (ts TeeSink) Complete(job string, status CompletionStatus, timing time.Duration, kv KV) {
	for _, sink := range ts {
		sink.Complete(job, status, timing, kv)
	}
}

type KVSink struct {
	KV   KV
	Sink Sink
}

func (s *KVSink) KeyValue(key string, val interface{}) {
	if s.KV == nil {
		s.KV = make(KV)
	}

	s.KV[key] = val
}

func (s *KVSink) mergedKV(kv KV) KV {
	if kv == nil {
		return s.KV
	}
	if s.KV == nil {
		return kv
	}

	merged := make(KV)

	for k, v := range s.KV {
		merged[k] = v
	}

	for k, v := range kv {
		merged[k] = v
	}

	return merged
}

func (s KVSink) Event(job, event string, kv KV) {
	s.Sink.Event(job, event, s.mergedKV(kv))
}

func (s KVSink) Error(job, event string, err error, kv KV) {
	s.Sink.Error(job, event, err, s.mergedKV(kv))
}

func (s KVSink) Timing(job, event string, timing time.Duration, kv KV) {
	s.Sink.Timing(job, event, timing, s.mergedKV(kv))
}

func (s KVSink) Gauge(job, event string, gauge float64, kv KV) {
	s.Sink.Gauge(job, event, gauge, s.mergedKV(kv))
}

func (s KVSink) Complete(job string, status CompletionStatus, timing time.Duration, kv KV) {
	s.Sink.Complete(job, status, timing, s.mergedKV(kv))
}
