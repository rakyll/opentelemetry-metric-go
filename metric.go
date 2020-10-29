package metric

import (
	"context"

	"go.opencensus.io/otel/api/metric"
)

type InstrumentOption = metric.InstrumentOption

type Measurement = metric.Measurement

type Meter struct{}

func (m Meter) NewFloat64Counter(name string, options ...InstrumentOption) (count.Float64Counter, error) {
}

func (m Meter) NewFloat64UpDownCounter(name string, options ...InstrumentOption) (count.Float64UpDownCounter, error) {
}

func (m Meter) NewInt64Counter(name string, options ...InstrumentOption) (count.Int64Counter, error) {
}

func (m Meter) NewInt64UpDownCounter(name string, options ...InstrumentOption) (count.Int64UpDownCounter, error) {
}

func (m Meter) NewFloat64SumObserver(name string, fn sum.Float64ObserverFunc, options ...InstrumentOption) (sum.Float64SumObserver, error) {
}

func (m Meter) NewInt64SumObserver(name string, fn sum.Int64ObserverFunc, options ...InstrumentOption) (sum.Int64SumObserver, error) {
}

func (m Meter) NewFloat64ValueRecorder(name string, options ...InstrumentOption) (sum.Float64SumObserver, error) {
}

func (m Meter) NewFloat64ValueObserver(name string, options ...InstrumentOption) (sum.Int64SumObserver, error) {
}

func (m Meter) NewInt64ValueRecorder(name string, options ...InstrumentOption) (sum.Float64SumObserver, error) {
}

func (m Meter) NewInt64ValueObserver(name string, options ...InstrumentOption) (sum.Int64SumObserver, error) {
}

func (m Meter) RecordBatch(ctx context.Context, ls []label.KeyValue, ms ...Measurement) {}

type BatchObserverResult = metric.BatchObserverResult

func (m Meter) NewBatchObserver(fn func(result BatchObserverResult)) {}
