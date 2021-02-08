package service

// services as application layer
// the application connects the model with repository layer

import (
	"context"
	"database/sql"

	"github.com/muhfaris/adhouse-sample/user/domain"
	"github.com/muhfaris/adhouse-sample/user/repository"
	"github.com/muhfaris/adhouse-sample/user/repository/psql"
)

// UserService is user service
type UserService struct {
	mutate repository.UserMutate
}

// NewUserService is to create RuleStatusService psql instance
func NewUserService(db *sql.DB) *UserService {
	loginService := &UserService{
		mutate: psql.NewUserMutateInPSQl(db),
	}

	return loginService
}

// AddUser is service login user
func (service *UserService) AddUser(ctx context.Context, username, password string) (domain.User, error) {
	user := domain.CreateUser(username, password)
	result := <-service.mutate.AddUser(ctx, *user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return result.Result.(domain.User), nil
}

// User is service login user
func (service *UserService) Login(ctx context.Context, username, password string) (domain.User, error) {
	result := <-service.mutate.Login(ctx, username, password)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return result.Result.(domain.User), nil
}
