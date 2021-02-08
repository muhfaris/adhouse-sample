package psql

import (
	"context"
	"database/sql"

	"github.com/muhfaris/adhouse-sample/user/domain"
	"github.com/muhfaris/adhouse-sample/user/repository"
)

// UserMutateInPSQL is implementation in postgres sql
type UserMutateInPSQL struct {
	Conn *sql.DB
}

// NewUserMutateInPSQl is create instance
func NewUserMutateInPSQl(db *sql.DB) repository.UserMutate {
	return &UserMutateInPSQL{Conn: db}
}

// AddUser is add new user
func (q *UserMutateInPSQL) AddUser(ctx context.Context, user domain.User) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)
	go func() {
		query := `
		INSERT(
			username,
			password
		)
		VALUES(
			username = $1,
			password = $2)`

		var loginModel domain.User
		_, err := q.Conn.ExecContext(ctx, query, user.Username, user.Password)
		if err != nil {
			result <- repository.QueryResult{Error: err}
			return
		}

		result <- repository.QueryResult{Result: loginModel}
	}()

	return result
}

// Login is check username and password
// sebaiknya password di rubah dengan hash  bcrypt, sha256 dll, jgn menggunakan md5.
// data (username, password) hanya sebagai contoh

func (q *UserMutateInPSQL) Login(ctx context.Context, username, password string) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)
	go func() {
		query := `
		SELECT
			username,
			password
		FROM
			users
		WHERE
			username = $1
		AND
			password = $2`

		var user domain.User
		err := q.Conn.QueryRowContext(ctx, query, username, password).Scan(
			&user.Username,
			&user.Password,
		)
		if err != nil {
			result <- repository.QueryResult{Error: err}
			return
		}

		result <- repository.QueryResult{Result: user}
	}()

	return result
}
