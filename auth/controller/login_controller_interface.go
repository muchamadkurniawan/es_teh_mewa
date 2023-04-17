package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LoginController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	LoginCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
