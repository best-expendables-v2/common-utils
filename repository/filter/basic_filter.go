package filter

import "strings"

type BasicFilter struct {
	where   Where
	joins   Joins
	keys    map[string]bool
	groups  Groups
	orWhere OrWhere
}

func NewBasicFilter() *BasicFilter {
	return &BasicFilter{
		where: Where{},
		joins: Joins{
			Queries: []string{},
			Conds:   make(map[string]Join),
		},
		keys:    Keys{},
		groups:  Groups{},
		orWhere: OrWhere{},
	}
}

func (f *BasicFilter) GetLimit() int {
	return IgnoreLimit
}

func (f *BasicFilter) GetOffset() int {
	return IgnoreOffset
}

func (f *BasicFilter) GetWhere() Where {
	return f.where
}

func (f *BasicFilter) GetOrWhere() OrWhere {
	return f.orWhere
}

func (f *BasicFilter) GetJoins() []Join {
	joins := []Join{}
	for _, join := range f.joins.Queries {
		joins = append(joins, f.joins.Conds[join])
	}
	return joins
}

func (f *BasicFilter) AddWhere(key string, query string, values ...interface{}) *BasicFilter {
	f.where[query] = values
	f.keys[key] = true
	return f
}

func (f *BasicFilter) AddOrWhere(key string, query string, values ...interface{}) *BasicFilter {
	f.where[query] = values
	f.keys[key] = true
	return f
}

func (f *BasicFilter) AddJoin(join string, values ...interface{}) *BasicFilter {
	_, exist := f.joins.Conds[join]
	if !exist {
		f.joins.Queries = append(f.joins.Queries, join)
	}

	f.joins.Conds[join] = Join{Query: join, Args: values}
	return f
}

func (f *BasicFilter) AddKey(key string) *BasicFilter {
	f.keys[key] = true
	return f
}

func (f *BasicFilter) GetKeys() Keys {
	return f.keys
}

func (f *BasicFilter) AddGroup(query string) *BasicFilter {
	f.groups[query] = true
	return f
}
func (f *BasicFilter) GetGroups() string {
	var queries []string
	for query := range f.groups {
		queries = append(queries, query)
	}
	return strings.Join(queries, ",")
}
