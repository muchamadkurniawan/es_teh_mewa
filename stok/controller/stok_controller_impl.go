package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/stok/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type StokControllerImpl struct {
	service service.StokService
}

func NewStokController(serv service.StokService) StokController {
	return &StokControllerImpl{
		service: serv,
	}
}

func (s *StokControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (s *StokControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := s.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"stok/view/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	serv := s.service.GetAll(context.Background())
	myTemplate.ExecuteTemplate(w, "indexStok", map[string]interface{}{
		"Title":   "Cafe Mewa - Stok",
		"Nama":    session["nama"],
		"AllData": serv,
	})
}

func (s *StokControllerImpl) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := s.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"stok/view/indexDetail.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	serv := s.service.GetAllDetail(context.Background())
	myTemplate.ExecuteTemplate(w, "indexDetailStok", map[string]interface{}{
		"Title":   "Cafe Mewa - Stok",
		"Nama":    session["nama"],
		"AllData": serv,
	})
}
