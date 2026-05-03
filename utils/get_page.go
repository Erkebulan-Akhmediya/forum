package utils

import (
	"net/http"
	"strconv"
)

const (
	defaultPageIndex = 0
	defaultPageSize  = 50
)

type Page struct {
	Index, Size int
}

// returns pagination option in Page struct
func GetPage(r *http.Request) *Page {
	var p Page
	var err error
	pageSize := r.URL.Query().Get("page_size")
	p.Size, err = strconv.Atoi(pageSize)
	if err != nil {
		p.Size = defaultPageSize
	}

	pageIndex := r.URL.Query().Get("page_index")
	p.Index, err = strconv.Atoi(pageIndex)
	if err != nil {
		p.Index = defaultPageIndex
	}

	return &p
}
