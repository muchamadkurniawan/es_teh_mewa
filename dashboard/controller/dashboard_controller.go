package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DashboardContrller interface {
	CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{}
	GetRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DetailRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateBiayaPesanan(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
