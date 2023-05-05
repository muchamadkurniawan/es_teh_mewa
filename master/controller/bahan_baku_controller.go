package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/master/service"
	"es_teh_mewa/master/web"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type BahanBakuControllerImpl struct {
	service service.BahanbakuService
}

func NewBahanBakuController(BahanBakuService service.BahanbakuService) BahanBakuController {
	return &BahanBakuControllerImpl{
		service: BahanBakuService,
	}
}

func (controller *BahanBakuControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *BahanBakuControllerImpl) FindAllBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	file := controller.service.FindAll(context.Background())
	myTemplate := template.Must(template.ParseFiles("master/views/bahan_baku/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "indexBahanBaku", map[string]interface{}{
		"Title":  "Cafe Mewa - Bahan Baku",
		"Nama":   session["nama"],
		"data":   file,
		"satuan": controller.service.GetSatuan(context.Background()),
	})
}

func (controller *BahanBakuControllerImpl) FindByIdBahanBaku(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	file := controller.service.FindById(context.Background(), id)
	satuan := controller.service.GetSatuan(context.Background())
	myTemplate := template.Must(template.ParseFiles("master/views/bahan_baku/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "showBahanBaku", map[string]interface{}{
		"Title":  "Cafe Mewa - Bahan Baku",
		"Nama":   session["nama"],
		"data":   file,
		"satuan": satuan,
	})
}

func (controller *BahanBakuControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("master/views/bahan_baku/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "createBahanBaku", map[string]interface{}{
		"Title": "Cafe Mewa - Bahan Baku",
		"Nama":  session["nama"],
	})
}

func (controller *BahanBakuControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	idSatuan, err := strconv.Atoi(r.PostFormValue("satuan"))
	helperMain.PanicIfError(err)
	nama := r.PostFormValue("nama")
	file := web.BahanbakuRequest{
		IdSatuan: idSatuan,
		Nama:     nama,
	}
	controller.service.Save(context.Background(), file)
	http.Redirect(w, r, "/bahan-baku/", http.StatusFound)
}

func (controller *BahanBakuControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	idSatuan, err := strconv.Atoi(r.PostFormValue("satuan"))
	helperMain.PanicIfError(err)
	nama := r.PostFormValue("nama")
	file := web.BahanbakuRequest{
		Id:       id,
		IdSatuan: idSatuan,
		Nama:     nama,
	}
	controller.service.Update(context.Background(), file)
	http.Redirect(w, r, "/bahan-baku/", http.StatusFound)
}

func (controller *BahanBakuControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	controller.service.Delete(context.Background(), id)
	http.Redirect(w, r, "/bahan-baku/", http.StatusFound)
}
