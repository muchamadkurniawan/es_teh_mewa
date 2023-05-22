package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BiayaController interface {
	CheckLoginKasir(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	CheckLoginAdmin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateBiaya(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByIdAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
