package logr

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// check that WriterSink implements the Sink interface
var _ Sink = DiscardSink{}

type WriterSink struct {
	w io.Writer
}

func formattedKV(kv KV) string {
	keys := make([]string, 0, len(kv))
	for k, _ := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	formatted := make([]string, 0, len(keys))
	for _, k := range keys {
		formatted = append(formatted, fmt.Sprintf("%s:%#v", k, kv[k]))
	}

	return strings.Join(formatted, " ")
}

func NewWriterSink(w io.Writer) WriterSink {
	return WriterSink{
		w: w,
	}
}

func (s WriterSink) Event(job, event string, kv KV) {
	fmt.Fprintf(s.w, "[%s] job:%s event:%s kv:[%s]\n",
		time.Now().Format(time.RFC3339Nano),
		job, event, formattedKV(kv),
	)
}

func (s WriterSink) Error(job, event string, err error, kv KV) {
	fmt.Fprintf(s.w, "[%s] job:%s event:%s error:%s kv:[%s]\n",
		time.Now().Format(time.RFC3339Nano),
		job, event, err.Error(), formattedKV(kv),
	)
}

func (s WriterSink) Timing(job, event string, timing time.Duration, kv KV) {
	fmt.Fprintf(s.w, "[%s] job:%s event:%s timing:%s kv:[%s]\n",
		time.Now().Format(time.RFC3339Nano),
		job, event, timing.String(), formattedKV(kv),
	)
}

func (s WriterSink) Gauge(job, event string, value float64, kv KV) {
	fmt.Fprintf(s.w, "[%s] job:%s event:%s gauge:%g kv:[%s]\n",
		time.Now().Format(time.RFC3339Nano),
		job, event, value, formattedKV(kv),
	)
}

func (s WriterSink) Complete(job string, status CompletionStatus, timing time.Duration, kv KV) {
	fmt.Fprintf(s.w, "[%s] job:%s status:%s, timing:%s kv:[%s]\n",
		time.Now().Format(time.RFC3339Nano),
		job, status, timing.String(), formattedKV(kv),
	)
}
