package filter

const (
	defaultPerPage = 50
	maxPerPage     = 200
)

type PaginationFilter struct {
	BasicFilter   `json:"basicFilter"`
	BasicOrder    `json:"basicOrder"`
	CheckNextPage bool `json:"checkNextPage"`
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
	if f.CheckNextPage {
		return f.GetPerPage() + 1
	}
	return f.GetPerPage()
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
