package metric

import (
	"context"

	"go.opencensus.io/otel/api/metric"
)

type InstrumentOption = metric.InstrumentOption

type Measurement = metric.Measurement

type Meter struct{}

func (m Meter) NewFloat64Counter(name string, options ...InstrumentOption) (count.Float64Counter, error) {
	panic("not eyet")
}

func (m Meter) NewFloat64UpDownCounter(name string, options ...InstrumentOption) (count.Float64UpDownCounter, error) {
	panic("not eyet")
}

func (m Meter) RecordBatch(ctx context.Context, ls []label.KeyValue, ms ...Measurement) {}

type BatchObserverResult struct{}

func (m Meter) NewBatchObserver(fn func(result BatchObserverResult)) {}
