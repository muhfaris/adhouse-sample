package handler

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/schema"
	response "github.com/muhfaris/adhouse-sample/helper/api_response"
	"github.com/muhfaris/adhouse-sample/product/service"
	"github.com/muhfaris/adhouse-sample/product/structures"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// ProductHandler is wrap user all handler
type productHandler struct {
	service *service.ProductService
	logger  *log.Logger
	decoder *schema.Decoder
}

// NewProductHandler is initialize user handler
func NewProductHandler(db *sql.DB, logger *logrus.Logger) *productHandler {
	return &productHandler{
		service: service.NewProductService(db),
		logger:  logger,
		decoder: schema.NewDecoder(),
	}
}

// GetProductByID is wrap Login handler
func (ph *productHandler) GetProductByID(w http.ResponseWriter, r *http.Request) response.Response {
	ctx := r.Context()
	paramQuery := r.URL.Query() // id=[1,2,3,"]&name=pc

	var product structures.ProductRead
	err := ph.decoder.Decode(&product, paramQuery)
	if err != nil {
		return response.Response{StatusCode: http.StatusNotAcceptable, Error: err, Message: err.Error()}
	}

	resp, err := ph.service.GetProductDetailByID(ctx, product)
	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Error: err, Message: err.Error()}
	}

	return response.Response{Data: resp, StatusCode: http.StatusOK}
}
