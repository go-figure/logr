package logr_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/go-figure/logr"
	"github.com/stretchr/testify/require"
)

type testJSONEntry struct {
	Timestamp string                `json:"timestamp"`
	Job       string                `json:"job"`
	Event     string                `json:"event"`
	Status    logr.CompletionStatus `json:"status,omitempty"`
	Error     string                `json:"error,omitempty"`
	Timing    time.Duration         `json:"timing,omitempty"`
	Gauge     float64               `json:"gauge,omitempty"`
	KV        logr.KV               `json:"kv"`
}

func TestJSONWriterSink(t *testing.T) {
	var buf bytes.Buffer
	sink := logr.NewJSONWriterSink(&buf)

	kv := logr.KV{
		"key1": 1.0,
		"key2": "val",
		"key3": true,
	}
	err := errors.New("testing error")
	timing := 42 * time.Second
	gauge := 0.42

	sink.Event("testing", "event", kv)
	checkJSONEntry(t, &buf, testJSONEntry{
		Job:   "testing",
		Event: "event",
		KV:    kv,
	})

	sink.Error("testing", "event", err, kv)
	checkJSONEntry(t, &buf, testJSONEntry{
		Job:   "testing",
		Event: "event",
		Error: err.Error(),
		KV:    kv,
	})

	sink.Timing("testing", "event", timing, kv)
	checkJSONEntry(t, &buf, testJSONEntry{
		Job:    "testing",
		Event:  "event",
		Timing: timing,
		KV:     kv,
	})

	sink.Gauge("testing", "event", gauge, kv)
	checkJSONEntry(t, &buf, testJSONEntry{
		Job:   "testing",
		Event: "event",
		Gauge: gauge,
		KV:    kv,
	})

	sink.Complete("testing", logr.Success, timing, kv)
	checkJSONEntry(t, &buf, testJSONEntry{
		Job:    "testing",
		Status: logr.Success,
		Timing: timing,
		KV:     kv,
	})
}

func checkJSONEntry(t *testing.T, buf *bytes.Buffer, entry testJSONEntry) {
	var rcvEntry testJSONEntry
	err := json.NewDecoder(buf).Decode(&rcvEntry)
	require.Nil(t, err, "entry should be valid json")

	entry.Timestamp = rcvEntry.Timestamp
	require.Equal(t, entry, rcvEntry, "entries should match")

	timestamp, err := time.Parse(time.RFC3339Nano, rcvEntry.Timestamp)
	require.Nil(t, err, "timestamp should be RFC3339Nano")
	require.WithinDuration(t, time.Now(), timestamp, time.Millisecond, "timestamp should be time.Now")

	buf.Reset()
}
