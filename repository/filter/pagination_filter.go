package filter

const (
	defaultPerPage = 50
	maxPerPage     = 200
)

type PaginationFilter struct {
	BasicFilter
	BasicOrder

	Page          int  `json:"page"`
	PerPage       int  `json:"perPage"`
	IgnorePerPage bool `json:"ignorePerPage"`
}

func NewPaginationFilter() *PaginationFilter {
	return &PaginationFilter{
		BasicFilter: *NewBasicFilter(),
		BasicOrder:  *NewBasicOrder(),
	}
}

func (f *PaginationFilter) GetLimit() int {
	return f.GetPerPage() + 1
}

func (f *PaginationFilter) GetOffset() int {
	return (f.GetPage() - 1) * f.GetPerPage()
}

func (f *PaginationFilter) GetPage() int {
	if f.Page < 1 {
		return 1
	}
	return f.Page
}

func (f *PaginationFilter) GetPerPage() int {
	if f.PerPage < 1 || (f.PerPage > maxPerPage && !f.IgnorePerPage) {
		return defaultPerPage
	}

	return f.PerPage
}
