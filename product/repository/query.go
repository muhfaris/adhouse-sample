package repository

import "context"

// ProductQuery is product query
type ProductQuery interface {
	GetProductByID(context.Context, []int) <-chan QueryResult
}
