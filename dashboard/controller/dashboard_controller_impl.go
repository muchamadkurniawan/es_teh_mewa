package controller

import (
	"eh_teh_mewa/dashboard/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DashboardControllerImpl struct {
	service service.DashboardServiceImpl
}

func (service *DashboardControllerImpl) Index(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
