package filter

type Paging struct {
	Total       int `json:"total"`
	PerPage     int `json:"perPage"`
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
	From        int `json:"from"`
	To          int `json:"to"`
}

func GetPaging(f Filter, count int) *Paging {
	var lastPage int
	var from = f.GetOffset() + 1
	lastPage = int(count) / f.GetLimit()
	if lastPage*f.GetLimit() != int(count) {
		lastPage = lastPage + 1
	}

	var to = f.GetOffset() + f.GetLimit()
	if to > int(count) {
		to = int(count)
	}

	if lastPage < f.GetPage() {
		lastPage = 0
		from = 0
		to = 0
	}
	return &Paging{
		Total:       int(count),
		PerPage:     f.GetLimit(),
		CurrentPage: f.GetPage(),
		LastPage:    lastPage,
		From:        from,
		To:          to,
	}
}
