package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
func (q *ProductQueryInPSQL) GetProductByID(ctx context.Context, IDs []int, name string) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	var conditions []string
	var counter int = 1
	if len(IDs) > 1 {
		conditions = append(conditions, fmt.Sprintf("id in (%s)", convertPlaceholderInt(IDs)))
	}

	if name != "" {
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", name))
		counter++
	}

	where := strings.Join(conditions, " OR ")

	go func() {
		query := fmt.Sprintf(`
		SELECT
			id,
			name,
			qty
		FROM
			products
		WHERE
			%s`, where)

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
