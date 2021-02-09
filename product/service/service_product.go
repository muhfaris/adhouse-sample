package service

import (
	"context"
	"database/sql"

	"github.com/muhfaris/adhouse-sample/product/domain"
	"github.com/muhfaris/adhouse-sample/product/repository"
	"github.com/muhfaris/adhouse-sample/product/repository/psql"
	"github.com/muhfaris/adhouse-sample/product/structures"
)

// ProductService is user service
type ProductService struct {
	query repository.ProductQuery
}

// NewProductService is to create RuleStatusService psql instance
func NewProductService(db *sql.DB) *ProductService {
	loginService := &ProductService{
		query: psql.NewProductQueryInPSQl(db),
	}

	return loginService
}

// Login is service login user
func (service *ProductService) GetProductDetailByID(ctx context.Context, product structures.ProductRead) ([]domain.Product, error) {
	result := <-service.query.GetProductByID(ctx, product.ID, product.Name)
	if result.Error != nil {
		return []domain.Product{}, result.Error
	}

	return result.Result.([]domain.Product), nil
}
