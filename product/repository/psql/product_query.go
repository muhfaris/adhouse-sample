package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/muhfaris/adhouse-sample/product/domain"
	"github.com/muhfaris/adhouse-sample/product/repository"
)

// ProductQueryInPSQL is implementation in postgres sql
type ProductQueryInPSQL struct {
	Conn *sql.DB
}

// NewProductQueryInPSQl is create instance
func NewProductQueryInPSQl(db *sql.DB) repository.ProductQuery {
	return &ProductQueryInPSQL{Conn: db}
}

// Login is check username and password
// sebaiknya password di rubah dengan hash  bcrypt, sha256 dll, jgn menggunakan md5.
// data (username, password) hanya sebagai contoh
func (q *ProductQueryInPSQL) GetProductByID(ctx context.Context, IDs []int) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	conditionQuery := convertPlaceholderInt(IDs)
	go func() {
		query := fmt.Sprintf(`
		SELECT
			id,
			name,
			qty
		FROM
			products
		WHERE
			id in %s`, conditionQuery)

		rows, err := q.Conn.QueryContext(ctx, query)
		if err != nil {
			result <- repository.QueryResult{Error: err}
			return
		}

		var products []domain.Product
		for rows.Next() {
			var product domain.Product
			err := rows.Scan(
				&product.ID,
				&product.Name,
				&product.QTY,
			)

			if err != nil {
				result <- repository.QueryResult{Error: err}
				return
			}

			products = append(products, product)
		}

		result <- repository.QueryResult{Result: products}
	}()

	return result
}
