package filter

import (
	"fmt"
	"strings"
)

type BasicOrder struct {
	OrderBy               []string `json:"orderBy"`
	mapObjectToFieldOrder map[string]string
}

func NewBasicOrder() *BasicOrder {
	return &BasicOrder{
		mapObjectToFieldOrder: map[string]string{},
	}
}

func (s *BasicOrder) GetOrderBy() []string {
	return s.OrderBy
}

func (s BasicOrder) SetOrderBy(mapObjectToFieldOrder map[string]string) []string {
	s.allowObjectToOrder(mapObjectToFieldOrder)
	orderByClauses := make([]string, 0)
	for i := 0; i < len(s.OrderBy); i++ {
		modifiedClause := s.assignTableToOrderByClause(s.OrderBy[i])
		if modifiedClause == "" {
			continue
		}
		orderByClauses = append(orderByClauses, modifiedClause)
	}
	return orderByClauses
}

func (s *BasicOrder) allowObjectToOrder(mapObjectToFieldOrder map[string]string) {
	s.mapObjectToFieldOrder = mapObjectToFieldOrder
}

func (s BasicOrder) assignTableToOrderByClause(orderByClause string) string {
	words := strings.Split(orderByClause, " ")
	if len(words) != 2 {
		return ""
	}
	for object, field := range s.mapObjectToFieldOrder {
		fmt.Println(object, words[0])
		if object == words[0] {
			return fmt.Sprintf("%s %s", field, words[1])
		}
	}
	return ""
}
