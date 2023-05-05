package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RekapController interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
