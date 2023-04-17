package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PembelianController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
