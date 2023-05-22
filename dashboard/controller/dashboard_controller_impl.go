package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/dashboard/service"
	"es_teh_mewa/helperMain"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type DashboardControllerImpl struct {
	service service.DashboardService
}

func NewDashboardController(service service.DashboardService) DashboardContrller {
	return &DashboardControllerImpl{
		service: service,
	}
}

func (controller *DashboardControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (d DashboardControllerImpl) GetRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := d.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"dashboard/view/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	data := d.service.GetRekap(context.Background())
	myTemplate.ExecuteTemplate(w, "indexDashboard", map[string]interface{}{
		"Title": "Cafe Mewa",
		"Nama":  session["nama"],
		"Data":  data,
	})
}

func (d *DashboardControllerImpl) DetailRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := d.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"dashboard/view/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	data := d.service.GetRekapById(context.Background(), id)
	pesanan := d.service.GetPesananRekapById(context.Background(), id)
	biaya := d.service.GetBiayaRekapById(context.Background(), id)
	myTemplate.ExecuteTemplate(w, "showDashboard", map[string]interface{}{
		"Title":   "Cafe Mewa",
		"Nama":    session["nama"],
		"Data":    data,
		"Pesanan": pesanan,
		"Biaya":   biaya,
	})
}
