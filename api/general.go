package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Health() httprouter.Handle  {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		successResponse(w, "{\"status\":\"healthy\"}")
	}
}

func successResponse(w http.ResponseWriter, v interface{})  {
	b, err := json.Marshal(&v)
	if err != nil {
		http500status(w)
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	_, err = w.Write(b)
	if err != nil {
		http500status(w)
	}
}

func http500status(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusInternalServerError)
}

func invalidRequest(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = io.WriteString(w, `{"code":400,"message":"invalid request body"}`)
}