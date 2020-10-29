package metric_test

import (
	"context"
	"log"

	metric "github.com/rakyll/opentelemetry-metric-go"
)

func Example_newFloat64Counter() {
	ctx := context.Background()

	var meter metric.Meter // ... get it from global provider
	counter, err := meter.AsFloat64().NewCounter("latency")
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 15.0)
}
