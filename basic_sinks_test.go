package logr_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go-figure/logr"
	"github.com/go-figure/logr/mock"
	"github.com/golang/mock/gomock"
)

func TestTeeSink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kv := logr.KV{
		"key1": 1,
		"key2": "val",
		"key3": true,
	}
	err := errors.New("testing error")
	timing := 42 * time.Second
	gauge := 0.42

	var ts logr.TeeSink
	for i := 0; i < 10; i++ {
		sink := mock.NewMockSink(ctrl)
		sink.EXPECT().Event("testing", "event", kv)
		sink.EXPECT().Error("testing", "event", err, kv)
		sink.EXPECT().Timing("testing", "event", timing, kv)
		sink.EXPECT().Gauge("testing", "event", gauge, kv)
		sink.EXPECT().Complete("testing", logr.Success, timing, kv)
		ts.Add(sink)
	}

	ts.Event("testing", "event", kv)
	ts.Error("testing", "event", err, kv)
	ts.Timing("testing", "event", timing, kv)
	ts.Gauge("testing", "event", gauge, kv)
	ts.Complete("testing", logr.Success, timing, kv)
}

func TestKVSink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kv := logr.KV{
		"key1": 1,
		"key2": "val",
		"key3": true,
	}
	err := errors.New("testing error")
	timing := 42 * time.Second
	gauge := 0.42

	sink := mock.NewMockSink(ctrl)
	sink.EXPECT().Event("testing", "event", kv)
	sink.EXPECT().Error("testing", "event", err, kv)
	sink.EXPECT().Timing("testing", "event", timing, kv)
	sink.EXPECT().Gauge("testing", "event", gauge, kv)
	sink.EXPECT().Complete("testing", logr.Success, timing, kv)

	kvs := logr.KVSink{
		Sink: sink,
	}
	kvs.KeyValue("key1", 1)
	kvs.KeyValue("key2", "val")

	kv = logr.KV{
		"key3": true,
	}

	kvs.Event("testing", "event", kv)
	kvs.Error("testing", "event", err, kv)
	kvs.Timing("testing", "event", timing, kv)
	kvs.Complete("testing", logr.Success, timing, kv)

	tKV := kvs.KV
	tKV["key3"] = true
	kvs.KV = nil
	kvs.Gauge("testing", "event", gauge, tKV)
}
