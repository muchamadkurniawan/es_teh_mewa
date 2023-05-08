package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/biaya/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type BiayaControllerImpl struct {
	service service.BiayaService
}

func NewBiayaController(serivce service.BiayaService) BiayaController {
	return &BiayaControllerImpl{
		service: serivce,
	}
}

func (controller *BiayaControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "kasir" {
		http.Redirect(w, r, "/user/", http.StatusFound)
	}
	return session.Values
}

func (controller *BiayaControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("biaya/view/kasir/index.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	bahan := controller.service.GetBahanBakuNonProdukServ(context.Background())
	today := controller.service.GetBiayaTodayServ(context.Background())
	myTemplate.ExecuteTemplate(w, "indexBiaya", map[string]interface{}{
		"Title": "Cafe Mewa - Biaya",
		"Nama":  session["nama"],
		"now":   time.Now().Format("02 Jan 2006"),
		"Bahan": bahan,
		"Today": today,
	})
}
