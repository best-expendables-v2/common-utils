package nrcontext

import (
	"fmt"
	"github.com/newrelic/go-agent"
	"net/http"
)

type TxnNameFunc func(*http.Request) string

type NewRelicMiddleware struct {
	app      newrelic.Application
	nameFunc TxnNameFunc
}

// Creates new middleware that will report time in NewRelic and set transaction in context
// It will ignore StatusNotFound, StatusBadRequest, StatusUnprocessableEntity
func NewMiddleware(appname string, license string) (*NewRelicMiddleware, error) {
	nrcfg := newrelic.NewConfig(appname, license)
	nrcfg.ErrorCollector.IgnoreStatusCodes = []int{
		http.StatusNotFound,
		http.StatusBadRequest,
		http.StatusUnprocessableEntity,
	}

	app, err := newrelic.NewApplication(nrcfg)

	if err != nil {
		return nil, err
	}
	return &NewRelicMiddleware{app, makeTransactionName}, nil
}

// Same as NewMiddleware but accepts newrelic.Config
func NewMiddlewareWithConfig(c newrelic.Config) (*NewRelicMiddleware, error) {
	app, err := newrelic.NewApplication(c)
	if err != nil {
		return nil, err
	}
	return &NewRelicMiddleware{app, makeTransactionName}, nil
}

// Same as NewMiddleware but accepts newrelic.Application
func NewMiddlewareWithApp(app newrelic.Application) *NewRelicMiddleware {
	return &NewRelicMiddleware{app, makeTransactionName}
}

// Allows to change transaction name. By default `fmt.Sprintf("%s %s", r.Method, r.URL.Path)`
func (nr *NewRelicMiddleware) SetTxnNameFunc(fn TxnNameFunc) {
	nr.nameFunc = fn
}

func makeTransactionName(r *http.Request) string {
	return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
}
