package nrredis

import (
	"github.com/go-redis/redis/v8"
	"github.com/newrelic/go-agent"
)

// WrapRedisClient adds newrelic measurements for commands and returns cloned client
func WrapRedisClient(txn newrelic.Transaction, c *redis.Client) *redis.Client {
	if txn == nil {
		return c
	}

	// clone using context
	ctx := c.Context()
	clone := c.WithContext(ctx)
	clone.AddHook(
		newRedisHook{
			txn: txn,
		},
	)
	return clone
}

type segment interface {
	End() error
}

// create segment through function to be able to test it
var segmentBuilder = func(txn newrelic.Transaction, product newrelic.DatastoreProduct, operation string) segment {
	return &newrelic.DatastoreSegment{
		StartTime: newrelic.StartSegmentNow(txn),
		Product:   product,
		Operation: operation,
	}
}
