package handler

import (
	"database/sql"
	"net/http"

	response "github.com/muhfaris/adhouse-sample/helper/api_response"
	pb "github.com/muhfaris/adhouse-sample/helper/parse_body"
	"github.com/muhfaris/adhouse-sample/user/service"
	"github.com/muhfaris/adhouse-sample/user/structures"
	log "github.com/sirupsen/logrus"
)

// UserHandler is wrap user all handler
type userHandler struct {
	service *service.UserService
	logger  *log.Logger
}

// NewUserHandler is initialize user handler
func NewUserHandler(db *sql.DB, logger *log.Logger) *userHandler {
	return &userHandler{
		service: service.NewUserService(db),
		logger:  logger,
	}
}

// AddUserHandler is wrap Login handler
func (ap *userHandler) AddUserHandler(w http.ResponseWriter, r *http.Request) response.Response {
	ctx := r.Context()
	var user structures.LoginRead
	err := pb.ParseBodyData(ctx, r, &user)
	if err != nil {
		return response.Response{StatusCode: http.StatusNotAcceptable, Error: err, Message: err.Error()}
	}

	resp, err := ap.service.AddUser(ctx, user.Username, user.Password)
	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Error: err, Message: err.Error()}
	}

	return response.Response{Data: resp, StatusCode: http.StatusOK}
}

// loginHandler is wrap Login handler
func (ap *userHandler) LoginHandler(w http.ResponseWriter, r *http.Request) response.Response {
	ctx := r.Context()
	var user structures.LoginRead
	err := pb.ParseBodyData(ctx, r, &user)
	if err != nil {
		return response.Response{StatusCode: http.StatusNotAcceptable, Error: err, Message: err.Error()}
	}

	resp, err := ap.service.Login(ctx, user.Username, user.Password)
	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Error: err, Message: err.Error()}
	}

	return response.Response{Data: resp, StatusCode: http.StatusOK}
}
