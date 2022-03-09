package connection

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type ConnPool struct {
	pluginRetry *pluginRetry
	gorm.ConnPool
}

func (s *ConnPool) String() string {
	return "gorm:db_retry:conn_pool"
}

func (s ConnPool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	result, err := s.ConnPool.PrepareContext(ctx, query)
	err = s.pluginRetry.retry(func() error {
		result, err = s.ConnPool.PrepareContext(ctx, query)
		return err
	}, err)
	return result, err
}

func (s ConnPool) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := s.ConnPool.ExecContext(ctx, query, args...)
	err = s.pluginRetry.retry(func() error {
		result, err = s.ConnPool.ExecContext(ctx, query, args...)
		return err
	}, err)
	return result, err
}

func (s ConnPool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	result, err := s.ConnPool.QueryContext(ctx, query, args...)
	err = s.pluginRetry.retry(func() error {
		result, err = s.ConnPool.QueryContext(ctx, query, args...)
		return err
	}, err)
	return result, err
}

func (s ConnPool) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.ConnPool.QueryRowContext(ctx, query, args...)
}

func (s *ConnPool) BeginTx(ctx context.Context, opt *sql.TxOptions) (gorm.ConnPool, error) {
	if basePool, ok := s.ConnPool.(gorm.ConnPoolBeginner); ok {
		result, err := basePool.BeginTx(ctx, opt)
		err = s.pluginRetry.retry(func() error {
			result, err = basePool.BeginTx(ctx, opt)
			return err
		}, err)
		return result, err
	}
	return s, nil
}

func (s *ConnPool) Commit() error {
	if basePool, ok := s.ConnPool.(gorm.TxCommitter); ok {
		err := basePool.Commit()
		err = s.pluginRetry.retry(func() error {
			err = basePool.Commit()
			return err
		}, err)
		return err
	}
	return nil
}

func (s *ConnPool) Rollback() error {
	if basePool, ok := s.ConnPool.(gorm.TxCommitter); ok {
		err := basePool.Rollback()
		err = s.pluginRetry.retry(func() error {
			err = basePool.Rollback()
			return err
		}, err)
		return err
	}
	return nil
}

func (s *ConnPool) Ping() error {
	return nil
}
