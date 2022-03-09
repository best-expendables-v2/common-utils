package connection

import (
	"github.com/best-expendables-v2/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type retryMessage struct {
	ErrMsgKey string
	Retry     int
	Delay     time.Duration
}

type RetryConfig struct {
	Attempt int           `envconfig:"GORM_RETRY_ATTEMPT" required:"true"`
	Delay   time.Duration `envconfig:"GORM_RETRY_DELAY" required:"true"`
}

type PluginRetry interface {
	Name() string
	Initialize(*gorm.DB) error
}

type pluginRetry struct {
	*gorm.DB
	ConnPool      *ConnPool
	retryConfig   RetryConfig
	retryMessages []retryMessage
}

func RegisterRetry() PluginRetry {
	return pluginRetry{}
}

func (s pluginRetry) Name() string {
	return "gorm:db_retry"
}

func (s pluginRetry) Initialize(db *gorm.DB) error {
	s.DB = db
	s.registerConnPool(db)
	s.retryConfig = s.loadConfig()
	s.retryMessages = s.getRetryConfig(s.retryConfig)
	return nil
}

func (s pluginRetry) loadConfig() RetryConfig {
	var config RetryConfig
	_ = godotenv.Load()
	envconfig.MustProcess("", &config)
	return config
}

func (s *pluginRetry) registerConnPool(db *gorm.DB) {
	basePool := db.ConnPool
	if _, ok := basePool.(ConnPool); ok {
		return
	}
	s.ConnPool = &ConnPool{ConnPool: basePool, pluginRetry: s}
	db.ConnPool = s.ConnPool
	db.Statement.ConnPool = s.ConnPool
}

func (s pluginRetry) getRetryConfig(conf RetryConfig) []retryMessage {
	var retryMessages []retryMessage
	retryMessages = append(retryMessages, retryMessage{"connection reset by peer", conf.Attempt, conf.Delay})
	retryMessages = append(retryMessages, retryMessage{"write: broken pipe", conf.Attempt, conf.Delay})
	retryMessages = append(retryMessages, retryMessage{"connection refused", conf.Attempt, conf.Delay})
	return retryMessages
}

func (s pluginRetry) retry(f func() error, err error) error {
	if err == nil {
		return nil
	}
	for _, retryConf := range s.retryMessages {
		if strings.Contains(err.Error(), retryConf.ErrMsgKey) {
			for i := 0; i < retryConf.Retry; i++ {
				err = f()
				if err == nil {
					return nil
				}
				logger.Error(errors.Wrap(err, "Retrying the execution"))
				time.Sleep(retryConf.Delay)
			}
		}
	}
	return err
}
