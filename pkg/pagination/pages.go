// Package pagination provides support for pagination requests and responses.
package pagination

import (
	"strconv"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 1
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 1000
)

// Pages represents a paginated list of data items.
type Pages struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
// The total parameter specifies the total number of data items.
// If total is less than 0, it means total is unknown.
func New(page, perPage, total int) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		// make sure pageCount cover the last unfull page
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
	}
}

// NewPage creates a Pages object using the query parameters found in the given HTTP request.
// count stands for the total number of items. Use -1 if this is unknown.
func NewPage(page, perPage string, count int) *Pages {
	pageInt := parseInt(page, 1)
	perPageInt := parseInt(perPage, DefaultPageSize)
	return New(pageInt, perPageInt, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *Pages) Limit() int {
	return p.PerPage
}