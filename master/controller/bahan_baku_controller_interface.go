package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BahanBakuController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	FindAllBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByIdBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Store(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
