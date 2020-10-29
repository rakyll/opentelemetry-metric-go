package metric

import (
	"context"

	"github.com/rakyll/opentelemetry-metric-go/counter"
	"go.opencensus.io/otel/api/metric"
)

type InstrumentOption = metric.InstrumentOption

// type Float64SumObserver = metric.Float64SumObserver
// type Float64ObserverResult = metric.Float64ObserverResult
// type Float64UpDownSumObserver = metric.Float64UpDownSumObserver
// type Float64ObserverFunc = metric.Float64ObserverFunc
// type Float64ValueObserver = metric.Float64ValueObserver
// type Float64ValueRecorder = metric.Float64ValueRecorder

type Measurement = metric.Measurement

type Meter struct{}

func (m Meter) NewFloat64Counter(name string, options ...InstrumentOption) (counter.Float64Counter, error) {
	panic("not eyet")
}

func (m Meter) NewFloat64UpDownCounter(name string, options ...InstrumentOption) (counter.Float64UpDownCounter, error) {
	panic("not eyet")
}

func (m Meter) RecordBatch(ctx context.Context, ls []label.KeyValue, ms ...Measurement) {}

type BatchObserverResult struct{}

func (m Meter) NewBatchObserver(fn func(result BatchObserverResult)) {}
