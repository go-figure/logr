package logr

import "time"

type CompletionStatus string

var (
	Success = CompletionStatus("success")
	Failed  = CompletionStatus("error")
	Invalid = CompletionStatus("invalid")
	Junk    = CompletionStatus("junk")
)

type KV map[string]interface{}

type Receiver interface {
	Event(event string, kv KV)
	Error(event string, err error, kv KV)
	Timing(event string, timing time.Duration, kv KV)
	Gauge(event string, gauge float64, kv KV)
}

type Sink interface {
	Event(job, event string, kv KV)
	Error(job, event string, err error, kv KV)
	Complete(job, event string, status CompletionStatus, timing time.Duration, kv KV)
	Timing(job, event string, timing time.Duration, kv KV)
	Gauge(job, event string, gauge float64, kv KV)
}
