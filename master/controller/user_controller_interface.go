package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UsersController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Create(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Store(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
