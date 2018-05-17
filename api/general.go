package api

import (
	"encoding/json"
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
	_, err = w.Write(b)
	if err != nil {
		http500status(w)
	}
}

func http500status(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusInternalServerError)
}