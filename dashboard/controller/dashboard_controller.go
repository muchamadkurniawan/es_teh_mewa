package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DashboardController interface {
	Index(w http.ResponseWriter, r http.Request, params httprouter.Params)
}
