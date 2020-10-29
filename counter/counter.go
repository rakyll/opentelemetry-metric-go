package counter

import "context"

type Int64Counter struct{}

func (c Int64Counter) Add(ctx context.Context, value float64, labels ...label.KeyValue) {}

func (c Int64Counter) Bind(labels ...label.KeyValue) Float64UpDownCounter {}

func (c Int64Counter) Measurement(value float64) measurement.Measurement {}

type Float64Counter struct{}

func (c Float64Counter) Add(ctx context.Context, value float64, labels ...label.KeyValue) {}

func (c Float64Counter) Bind(labels ...label.KeyValue) Float64UpDownCounter {}

func (c Float64Counter) Measurement(value float64) measurement.Measurement {}

type Float64UpDownCounter struct{}

func (c Float64UpDownCounter) Add(ctx context.Context, value float64, labels ...label.KeyValue) {}

func (c Float64UpDownCounter) Bind(labels ...label.KeyValue) Float64UpDownCounter {}

func (c Float64UpDownCounter) Measurement(value float64) measurement.Measurement {}
