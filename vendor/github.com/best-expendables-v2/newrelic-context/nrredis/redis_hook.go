package nrredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	newrelic "github.com/newrelic/go-agent"
	"strings"
)

type newRedisHook struct {
	txn newrelic.Transaction
}

func (h newRedisHook) BeforeProcess(ctx context.Context, _ redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h newRedisHook) AfterProcess(_ context.Context, cmd redis.Cmder) error {
	segmentBuilder(h.txn, newrelic.DatastoreRedis, strings.Split(cmd.Name(), " ")[0]).End()
	return nil
}

func (h newRedisHook) BeforeProcessPipeline(ctx context.Context, _ []redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h newRedisHook) AfterProcessPipeline(_ context.Context, _ []redis.Cmder) error {
	return nil
}
