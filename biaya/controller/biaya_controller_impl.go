package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/biaya/service"
	"es_teh_mewa/biaya/web"
	"es_teh_mewa/helperMain"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
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

func (controller *BiayaControllerImpl) CheckLoginKasir(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "kasir" {
		http.Redirect(w, r, "/user/", http.StatusFound)
	}
	return session.Values
}

func (controller *BiayaControllerImpl) CheckLoginAdmin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *BiayaControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
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

func (controller *BiayaControllerImpl) CreateBiaya(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLoginKasir(w, r)
	barang, _ := strconv.Atoi(r.PostFormValue("barang"))
	jumlah, _ := strconv.Atoi(r.PostFormValue("jumlah"))
	harga, _ := strconv.Atoi(r.PostFormValue("harga"))
	request := web.BiayaRequestCreate{
		Barang:      barang,
		Jumlah:      jumlah,
		HargaSatuan: harga,
		Total:       harga * jumlah,
	}
	err := controller.service.CreateBiaya(context.Background(), request)
	controller.service.UpdateStok(context.Background(), barang, jumlah)
	helperMain.PanicIfError(err)
	http.Redirect(w, r, "/biaya-kasir/", http.StatusFound)
}

func (controller *BiayaControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
	myTemplate := template.Must(template.ParseFiles("biaya/view/kasir/show.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	id, _ := strconv.Atoi(params.ByName("id"))
	data := controller.service.FindById(context.Background(), id)
	bahan := controller.service.GetBahanBakuNonProdukServ(context.Background())
	myTemplate.ExecuteTemplate(w, "showBiaya", map[string]interface{}{
		"Title": "Cafe Mewa - Detail Biaya",
		"Nama":  session["nama"],
		"now":   time.Now().Format("02 Jan 2006"),
		"Data":  data,
		"Bahan": bahan,
	})
}

func (controller *BiayaControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLoginKasir(w, r)
	id, _ := strconv.Atoi(params.ByName("id"))
	err := controller.service.Delete(context.Background(), id)
	helperMain.PanicIfError(err)
	http.Redirect(w, r, "/biaya-kasir/", http.StatusFound)
}

func (controller *BiayaControllerImpl) FindByIdAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginAdmin(w, r)
	myTemplate := template.Must(template.ParseFiles("biaya/view/kasir/showAdmin.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, _ := strconv.Atoi(params.ByName("id"))
	data := controller.service.FindById(context.Background(), id)
	bahan := controller.service.GetBahanBaku(context.Background())
	myTemplate.ExecuteTemplate(w, "showBiayaAdmin", map[string]interface{}{
		"Title": "Cafe Mewa - Detail Biaya",
		"Nama":  session["nama"],
		"now":   time.Now().Format("02 Jan 2006"),
		"Data":  data,
		"Bahan": bahan,
	})
}
