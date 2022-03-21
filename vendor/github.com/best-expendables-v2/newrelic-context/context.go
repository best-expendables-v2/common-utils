package nrcontext

import (
	"context"

	"github.com/best-expendables-v2/newrelic-context/nrgorm"
	"github.com/best-expendables-v2/newrelic-context/nrredis"
	"github.com/go-redis/redis/v8"
	newrelic "github.com/newrelic/go-agent"
	"gorm.io/gorm"
)

// Sets transaction from Context to gorm settings, returns cloned DB
func SetTxnToGorm(ctx context.Context, db *gorm.DB) *gorm.DB {
	txn := newrelic.FromContext(ctx)
	return nrgorm.SetTxnToGorm(txn, db)
}

// Gets transaction from Context and applies RedisWrapper, returns cloned client
func WrapRedisClient(ctx context.Context, c *redis.Client) *redis.Client {
	txn := newrelic.FromContext(ctx)
	return nrredis.WrapRedisClient(txn, c)
}
