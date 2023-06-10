package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type StokController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	CheckLoginKasir(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	IndexKasir(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
