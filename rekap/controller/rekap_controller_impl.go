package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/rekap/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type RekapControllerImpl struct {
	service service.RekapService
}

func NewRekapController(service service.RekapService) RekapController {
	return &RekapControllerImpl{
		service: service,
	}
}

func (controller *RekapControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "kasir" {
		http.Redirect(w, r, "/user/", http.StatusFound)
	}
	return session.Values
}

func (controller *RekapControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("rekap/views/kasir/index.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	data := controller.service.GetAllPesananToday(context.Background())
	currentTime := time.Now()
	myTemplate.ExecuteTemplate(w, "indexRekapKasir", map[string]interface{}{
		"Title": "Cafe Mawa - Rekap",
		"Nama":  session["nama"],
		"Data":  data,
		"now":   currentTime.Format("02 January 2006"),
	})
}
