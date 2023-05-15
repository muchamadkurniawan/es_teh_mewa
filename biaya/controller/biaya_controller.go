package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BiayaController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateBiaya(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
