package repository

import (
	"context"

	"github.com/muhfaris/adhouse-sample/user/domain"
)

// UserMutate is query of login
type UserMutate interface {
	Login(context.Context, string, string) <-chan QueryResult
	AddUser(context.Context, domain.User) <-chan QueryResult
}
