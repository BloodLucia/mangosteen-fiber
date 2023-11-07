package page

type Pager struct {
	Page    int `json:"page"`
	PerPage int `json:"per_Page"`
	Count   int `json:"count"`
}

type Page[T any] struct {
	Records []T `json:"records"`
	*Pager  `json:"pager"`
}

func Build[T any](list []T, pager *Pager) *Page[T] {
	return &Page[T]{
		Records: list,
		Pager:   pager,
	}
}
