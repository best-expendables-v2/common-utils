package repository

import (
	"context"
	"errors"

	"github.com/best-expendables-v2/common-utils/repository/filter"

	"github.com/best-expendables-v2/common-utils/model"
)

var RecordNotFound = errors.New("record not found")
var TransitionNotApplicable = errors.New("cannot transition to the given status")

type BaseRepo interface {
	Searchable
	Updatable
	Saveable
	Creatable
	Removable
	CanFindByID
	CanCreateOrUpdate
}

type PreloadField struct {
	FieldName  string
	Conditions []interface{}
}

func NewPreloadField(fieldName string, conditions ...interface{}) PreloadField {
	return PreloadField{
		FieldName:  fieldName,
		Conditions: conditions,
	}
}

type Searchable interface {
	Search(ctx context.Context, val interface{}, f filter.Filter, preloadFields ...string) error
	SearchWithPreloadCondition(ctx context.Context, val interface{}, f filter.Filter, preloadFields ...PreloadField) error
	SearchAndCount(ctx context.Context, val interface{}, f filter.Filter, preloadFields ...string) (int64, error)
}

type Updatable interface {
	Update(ctx context.Context, m model.Model, attrs ...interface{}) error
	Updates(ctx context.Context, m model.Model, params interface{}) error
}

type Saveable interface {
	Save(ctx context.Context, m model.Model) error
}

type Creatable interface {
	Create(ctx context.Context, m model.Model) error
}

type Removable interface {
	DeleteByID(ctx context.Context, m model.Model, id string) error
}

type CanFindByID interface {
	FindByID(ctx context.Context, m model.Model, id string, preloadFields ...string) error
	FindByIDWithPreloadCondition(ctx context.Context, m model.Model, id string, preloadFields ...PreloadField) error
}

type CanCreateOrUpdate interface {
	CreateOrUpdate(ctx context.Context, m model.Model, query interface{}, attrs ...interface{}) error
}
