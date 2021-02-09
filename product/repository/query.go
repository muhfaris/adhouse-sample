package repository

import "context"

// ProductQuery is product query
type ProductQuery interface {
	GetProductByID(context.Context, []int, string) <-chan QueryResult
}
