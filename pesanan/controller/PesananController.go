package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PesananController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Show(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Store(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Cetak(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	AddBiaya(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
