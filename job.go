package logr

import "time"

// check that Job implementes the Receiver interface
var _ Receiver = (*Job)(nil)

var DiscardReceiver = NewJob(Discard, "discard", nil)

type Job struct {
	Name      string
	StartedAt time.Time
	KVSink
}

func NewJob(sink Sink, name string, kv KV) *Job {
	if sink == nil {
		sink = Discard
	}

	job := &Job{
		Name:      name,
		StartedAt: time.Now(),
		KVSink: KVSink{
			KV:   kv,
			Sink: sink,
		},
	}

	job.Event("job_start", nil)
	return job
}

func (j *Job) Event(event string, kv KV) {
	j.KVSink.Event(j.Name, event, kv)
}

func (j *Job) Error(event string, err error, kv KV) {
	j.KVSink.Error(j.Name, event, err, kv)
}

func (j *Job) Complete(event string, status CompletionStatus, kv KV) {
	j.KVSink.Complete(j.Name, event, status, time.Since(j.StartedAt), kv)
}

func (j *Job) Timing(event string, timing time.Duration, kv KV) {
	j.KVSink.Timing(j.Name, event, timing, kv)
}

func (j *Job) Gauge(event string, gauge float64, kv KV) {
	j.KVSink.Gauge(j.Name, event, gauge, kv)
}
