package metric

import (
	"context"

	"github.com/rakyll/opentelemetry-metric-go/value"
	"go.opencensus.io/otel/api/metric"
)

type InstrumentOption = metric.InstrumentOption

type Measurement = metric.Measurement

type Meter struct{}

func (m Meter) AsInt64() Int64Meter {}

func (m Meter) AsFloat64() Int64Meter {}

type Float64Meter struct{}

func (m Float64Meter) NewCounter(name string, options ...InstrumentOption) (count.Float64Counter, error) {
}

func (m Float64Meter) NewUpDownCounter(name string, options ...InstrumentOption) (count.Float64UpDownCounter, error) {
}

func (m Float64Meter) NewSumObserver(name string, fn sum.Float64ObserverFunc, options ...InstrumentOption) (sum.Float64SumObserver, error) {
}

func (m Float64Meter) NewValueRecorder(name string, options ...InstrumentOption) (value.Float64Recorder, error) {
}

func (m Float64Meter) NewValueObserver(name string, options ...InstrumentOption) (value.Float64Observer, error) {
}

type Int64Meter struct{}

func (m Int64Meter) NewCounter(name string, options ...InstrumentOption) (count.Int64Counter, error) {
}

func (m Int64Meter) NewUpDownCounter(name string, options ...InstrumentOption) (count.Int64UpDownCounter, error) {
}

func (m Int64Meter) NewSumObserver(name string, fn sum.Int64ObserverFunc, options ...InstrumentOption) (sum.Int64SumObserver, error) {
}

func (m Int64Meter) NewValueRecorder(name string, options ...InstrumentOption) (sum.Float64SumObserver, error) {
}

func (m Int64Meter) NewValueObserver(name string, options ...InstrumentOption) (sum.Int64SumObserver, error) {
}

func (m Meter) RecordBatch(ctx context.Context, ls []label.KeyValue, ms ...Measurement) {}

type BatchObserverResult = metric.BatchObserverResult

func (m Meter) NewBatchObserver(fn func(result BatchObserverResult)) {}
