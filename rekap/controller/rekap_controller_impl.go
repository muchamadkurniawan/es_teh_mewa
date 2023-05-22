package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/helperMain"
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
	pesanan, totalPesanan := controller.service.GetAllPesananToday(context.Background())
	biaya, totalBiaya := controller.service.GetAllBiayaToday(context.Background())
	currentTime := time.Now()
	myTemplate.ExecuteTemplate(w, "indexRekapKasir", map[string]interface{}{
		"Title":        "Cafe Mawa - Rekap",
		"Nama":         session["nama"],
		"Pesanan":      pesanan,
		"TotalPesanan": totalPesanan,
		"Biaya":        biaya,
		"TotalBiaya":   totalBiaya,
		"now":          currentTime.Format("02 January 2006"),
	})
}

func (controller *RekapControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	//fmt.Fprintln(w, r.Form["pesanan"])
	//fmt.Fprintln(w, r.Form["biaya"])
	id, err := controller.service.Create(context.Background(), r.Form["keterangan"][0])
	helperMain.PanicIfError(err)
	err = controller.service.UpdateRekapPesananBiaya(context.Background(), id, r.Form["pesanan"], r.Form["biaya"])
	http.Redirect(w, r, "/rekap-kasir/", http.StatusFound)
}

func (controller *RekapControllerImpl) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
