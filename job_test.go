package logr_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go-figure/logr"
	"github.com/go-figure/logr/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kvs := logr.KV{"key": 1,
		"key2": "val",
		"key3": true,
	}
	err := errors.New("testing error")
	timing := 42 * time.Second
	gauge := 0.42

	sink := mock.NewMockSink(ctrl)
	sink.EXPECT().Event("testing", "job_start", nil)
	sink.EXPECT().Event("testing", "event", kvs)
	sink.EXPECT().Error("testing", "event", err, kvs)
	sink.EXPECT().Timing("testing", "event", timing, kvs)
	sink.EXPECT().Gauge("testing", "event", gauge, kvs)
	sink.EXPECT().Complete("testing", logr.Success, gomock.Any(), kvs)

	job := logr.NewJob(sink, "testing", nil)
	job.KV = logr.KV{
		"key": 1,
	}
	job.KeyValue("key2", "val")

	kvs = logr.KV{
		"key3": true,
	}

	job.Event("event", kvs)
	job.Error("event", err, kvs)
	job.Timing("event", timing, kvs)
	job.Gauge("event", gauge, kvs)
	job.Complete(logr.Success, kvs)

	job = logr.NewJob(nil, "testing", nil)
	require.Equal(t, logr.Discard, job.Sink, "providing nil as a sink to NewJob should set the Sink to Discard")
}
