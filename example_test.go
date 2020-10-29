package metric_test

import (
	"context"
	"log"

	metric "github.com/rakyll/opentelemetry-metric-go"
)

func Example() {
	ctx := context.Background()

	var meter metric.Meter // ...
	counter, err := meter.NewFloat64Counter("latency")
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 15.0)
}
