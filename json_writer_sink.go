package logr

import (
	"encoding/json"
	"io"
	"time"
)

// check that JSONWriterSink implements the Sink interface
var _ Sink = JSONWriterSink{}

type jsonEntry struct {
	Timestamp string           `json:"timestamp"`
	Job       string           `json:"job"`
	Event     string           `json:"event"`
	Status    CompletionStatus `json:"status,omitempty"`
	Error     string           `json:"error,omitempty"`
	Timing    time.Duration    `json:"timing,omitempty"`
	Gauge     float64          `json:"gauge,omitempty"`
	KV        KV               `json:"kv"`
}

type JSONWriterSink struct {
	enc *json.Encoder
}

func NewJSONWriterSink(w io.Writer) JSONWriterSink {
	return JSONWriterSink{
		enc: json.NewEncoder(w),
	}
}

func (s JSONWriterSink) encode(e *jsonEntry) {
	e.Timestamp = time.Now().Format(time.RFC3339Nano)
	s.enc.Encode(e)
}

func (s JSONWriterSink) Event(job, event string, kv KV) {
	s.encode(&jsonEntry{Job: job, Event: event, KV: kv})
}

func (s JSONWriterSink) Error(job, event string, err error, kv KV) {
	s.encode(&jsonEntry{Job: job, Event: event, Error: err.Error(), KV: kv})
}

func (s JSONWriterSink) Complete(job, event string, status CompletionStatus, timing time.Duration, kv KV) {
	s.encode(&jsonEntry{Job: job, Event: event, Status: status, Timing: timing, KV: kv})
}

func (s JSONWriterSink) Timing(job, event string, timing time.Duration, kv KV) {
	s.encode(&jsonEntry{Job: job, Event: event, Timing: timing, KV: kv})
}

func (s JSONWriterSink) Gauge(job, event string, gauge float64, kv KV) {
	s.encode(&jsonEntry{Job: job, Event: event, Gauge: gauge, KV: kv})
}
