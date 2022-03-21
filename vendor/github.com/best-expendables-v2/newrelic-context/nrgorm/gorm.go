package nrgorm

import (
	"fmt"
	"strings"

	"github.com/newrelic/go-agent"
	"gorm.io/gorm"
)

const (
	txnGormKey   = "newrelicTransaction"
	startTimeKey = "newrelicStartTime"
)

// SetTxnToGorm sets transaction to gorm settings, returns cloned DB
func SetTxnToGorm(txn newrelic.Transaction, db *gorm.DB) *gorm.DB {
	if txn == nil {
		return db
	}
	return db.Set(txnGormKey, txn)
}

// AddGormCallbacks adds callbacks to NewRelic, you should call SetTxnToGorm to make them work
func AddGormCallbacks(db *gorm.DB) {
	dialect := db.Statement.Dialector.Name()
	var product newrelic.DatastoreProduct
	switch dialect {
	case "postgres":
		product = newrelic.DatastorePostgres
	case "mysql":
		product = newrelic.DatastoreMySQL
	case "sqlite3":
		product = newrelic.DatastoreSQLite
	case "mssql":
		product = newrelic.DatastoreMSSQL
	default:
		return
	}
	callbacks := newCallbacks(product)
	registerCallbacks(db, "create", callbacks)
	registerCallbacks(db, "query", callbacks)
	registerCallbacks(db, "update", callbacks)
	registerCallbacks(db, "delete", callbacks)
	registerCallbacks(db, "row_query", callbacks)
}

type callbacks struct {
	product newrelic.DatastoreProduct
}

func newCallbacks(product newrelic.DatastoreProduct) *callbacks {
	return &callbacks{product}
}

func (c *callbacks) beforeCreate(db *gorm.DB)   { c.before(db.Statement) }
func (c *callbacks) afterCreate(db *gorm.DB)    { c.after(db.Statement, "INSERT") }
func (c *callbacks) beforeQuery(db *gorm.DB)    { c.before(db.Statement) }
func (c *callbacks) afterQuery(db *gorm.DB)     { c.after(db.Statement, "SELECT") }
func (c *callbacks) beforeUpdate(db *gorm.DB)   { c.before(db.Statement) }
func (c *callbacks) afterUpdate(db *gorm.DB)    { c.after(db.Statement, "UPDATE") }
func (c *callbacks) beforeDelete(db *gorm.DB)   { c.before(db.Statement) }
func (c *callbacks) afterDelete(db *gorm.DB)    { c.after(db.Statement, "DELETE") }
func (c *callbacks) beforeRowQuery(db *gorm.DB) { c.before(db.Statement) }
func (c *callbacks) afterRowQuery(db *gorm.DB)  { c.after(db.Statement, "") }

func (c *callbacks) before(s *gorm.Statement) {
	txn, ok := s.Get(txnGormKey)
	if !ok {
		return
	}
	s.Set(startTimeKey, newrelic.StartSegmentNow(txn.(newrelic.Transaction)))
}

func (c *callbacks) after(s *gorm.Statement, operation string) {
	startTime, ok := s.Get(startTimeKey)
	if !ok {
		return
	}
	if operation == "" {
		operation = strings.ToUpper(strings.Split(s.SQL.String(), " ")[0])
	}
	_ = segmentBuilder(
		startTime.(newrelic.SegmentStartTime),
		c.product,
		s.SQL.String(),
		operation,
		s.Table,
	).End()
}

func registerCallbacks(db *gorm.DB, name string, c *callbacks) {
	beforeName := fmt.Sprintf("newrelic:%v_before", name)
	afterName := fmt.Sprintf("newrelic:%v_after", name)
	gormCallbackName := fmt.Sprintf("gorm:%v", name)
	// gorm does some magic, if you pass CallbackProcessor here - nothing works
	switch name {
	case "create":
		db.Callback().Create().Before(gormCallbackName).Register(beforeName, c.beforeCreate)
		db.Callback().Create().After(gormCallbackName).Register(afterName, c.afterCreate)
	case "query":
		db.Callback().Query().Before(gormCallbackName).Register(beforeName, c.beforeQuery)
		db.Callback().Query().After(gormCallbackName).Register(afterName, c.afterQuery)
	case "update":
		db.Callback().Update().Before(gormCallbackName).Register(beforeName, c.beforeUpdate)
		db.Callback().Update().After(gormCallbackName).Register(afterName, c.afterUpdate)
	case "delete":
		db.Callback().Delete().Before(gormCallbackName).Register(beforeName, c.beforeDelete)
		db.Callback().Delete().After(gormCallbackName).Register(afterName, c.afterDelete)
	case "row":
		db.Callback().Row().Before(gormCallbackName).Register(beforeName, c.beforeRowQuery)
		db.Callback().Row().After(gormCallbackName).Register(afterName, c.afterRowQuery)
	}
}

type segment interface {
	End() error
}

// create segment through function to be able to test it
var segmentBuilder = func(
	startTime newrelic.SegmentStartTime,
	product newrelic.DatastoreProduct,
	query string,
	operation string,
	collection string,
) segment {
	return &newrelic.DatastoreSegment{
		StartTime:          startTime,
		Product:            product,
		ParameterizedQuery: query,
		Operation:          operation,
		Collection:         collection,
	}
}
