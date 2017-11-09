package logr

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testWriterEntry struct {
	Timestamp string
	Job       string
	Event     string
	Status    CompletionStatus
	Error     string
	Timing    time.Duration
	Gauge     float64
	KV        KV
}

func TestWriterSink(t *testing.T) {
	var buf bytes.Buffer
	sink := NewWriterSink(&buf)

	kv := KV{
		"key1": 1,
		"key2": "val",
		"key3": true,
	}
	err := errors.New("testing error")
	timing := 42 * time.Second
	gauge := 0.42

	sink.Event("testing", "event", kv)
	checkEntry(t, &buf, testWriterEntry{
		Job:   "testing",
		Event: "event",
		KV:    kv,
	})

	sink.Error("testing", "event", err, kv)
	checkEntry(t, &buf, testWriterEntry{
		Job:   "testing",
		Event: "event",
		Error: err.Error(),
		KV:    kv,
	})

	sink.Timing("testing", "event", timing, kv)
	checkEntry(t, &buf, testWriterEntry{
		Job:    "testing",
		Event:  "event",
		Timing: timing,
		KV:     kv,
	})

	sink.Gauge("testing", "event", gauge, kv)
	checkEntry(t, &buf, testWriterEntry{
		Job:   "testing",
		Event: "event",
		Gauge: gauge,
		KV:    kv,
	})

	sink.Complete("testing", Success, timing, kv)
	checkEntry(t, &buf, testWriterEntry{
		Job:    "testing",
		Status: Success,
		Timing: timing,
		KV:     kv,
	})
}

func checkEntry(t *testing.T, buf *bytes.Buffer, entry testWriterEntry) {
	line := buf.String()

	timestampStr := strings.Trim(strings.Split(line, " ")[0], "[]")
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	require.Nil(t, err, "timestamp should be RFC3339Nano")
	require.WithinDuration(t, time.Now(), timestamp, time.Millisecond, "timestamp should be time.Now")

	require.Contains(t, line, "job:"+entry.Job)

	if len(entry.Event) > 0 {
		require.Contains(t, line, "event:"+entry.Event)
	}

	if len(entry.Status) > 0 {
		require.Contains(t, line, "status:"+entry.Status)
	}

	if len(entry.Error) > 0 {
		require.Contains(t, line, "error:"+entry.Error)
	}

	if entry.Timing != 0 {
		require.Contains(t, line, "timing:"+entry.Timing.String())
	}

	if entry.Gauge != 0 {
		require.Contains(t, line, "gauge:"+strconv.FormatFloat(entry.Gauge, 'g', -1, 64))
	}

	require.Contains(t, line, "kv:["+formattedKV(entry.KV)+"]")

	buf.Reset()
}
