package metric_test

import (
	"context"
	"log"

	metric "github.com/rakyll/opentelemetry-metric-go"
)

func Example_asInt64() {
	ctx := context.Background()

	var meter metric.Meter // ... get it from global provider
	counter, err := meter.AsInt64().NewCounter("transactions")
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 15)
}
