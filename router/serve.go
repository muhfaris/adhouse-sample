package router

import (
	"encoding/json"
	"net/http"

	response "github.com/muhfaris/adhouse-sample/helper/api_response"
	log "github.com/sirupsen/logrus"
)

// customServe is wrap app handler
type customServe struct {
	*log.Logger
	HandlerFunc func(http.ResponseWriter, *http.Request) response.Response
}

// ServeHTTP is custom serve listen app
func (cs customServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	data := cs.HandlerFunc(w, r)
	if data.Error != nil {
		w.WriteHeader(data.StatusCode)
		cs.serveResponse(w, r, data.Error)
		return
	}

	w.WriteHeader(data.StatusCode)
	cs.serveResponse(w, r, data)
}

func (cs customServe) serveResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		cs.WithError(err).Println("encode error response")
		return
	}
}
