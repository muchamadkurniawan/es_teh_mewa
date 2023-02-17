package controller

import (
	"context"
	"eh_teh_mewa/master/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BahanBakuControllerImpl struct {
	service service.BahanbakuService
}

func NewBahanBakuController(service service.BahanbakuService) BahanBakuController {
	return &BahanBakuControllerImpl{
		service: service,
	}
}

func (controller *BahanBakuControllerImpl) FindAllBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	file := controller.service.FindAll(context.Background())
}

func (controller *BahanBakuControllerImpl) FindByIdBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *BahanBakuControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *BahanBakuControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *BahanBakuControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *BahanBakuControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
