package filter

const (
	IgnoreLimit  = -1
	IgnoreOffset = -1
)

type Where map[string][]interface{}
type Joins struct {
	Queries []string
	Conds   map[string]Join
}
type Join struct {
	Query string
	Args  []interface{}
}
type Keys map[string]bool
type Groups map[string]bool
type Filter interface {
	GetWhere() Where
	GetOrWhere() Where
	GetJoins() []Join
	GetLimit() int
	GetOffset() int
	GetOrderBy() []string
	GetKeys() Keys
	GetGroups() string
}
