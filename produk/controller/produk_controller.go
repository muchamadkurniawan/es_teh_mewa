package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/produk/model/entity"
	"es_teh_mewa/produk/model/web"
	"es_teh_mewa/produk/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type ProdukControllerImpl struct {
	service service.ProdukService
}

func NewProdukController(serviceProduk service.ProdukService) ProdukController {
	return &ProdukControllerImpl{
		service: serviceProduk,
	}
}

func (controller *ProdukControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *ProdukControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	file, err := controller.service.FindById(context.Background(), id)
	helperMain.PanicIfError(err)
	barang, _ := controller.service.GetBahan(context.Background())
	myTemplate := template.Must(template.ParseFiles("produk/views/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "showProduk", map[string]interface{}{
		"Title":  "Cafe Mewa - Produk",
		"Nama":   session["nama"],
		"data":   file,
		"barang": barang,
	})
}

func (controller *ProdukControllerImpl) FindByAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	var file []web.ResponseProdukFull
	var err error
	idBarang, err := strconv.Atoi(r.PostFormValue("barang_filter"))
	if r.PostFormValue("barang_filter") != "" {
		file, err = controller.service.FindAllByBahan(context.Background(), idBarang)
		helperMain.PanicIfError(err)
	} else {
		file, err = controller.service.FindByAll(context.Background())
		helperMain.PanicIfError(err)
	}
	barang, _ := controller.service.GetBahan(context.Background())
	myTemplate := template.Must(template.ParseFiles("produk/views/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "indexProduk", map[string]interface{}{
		"Title":      "Cafe Mewa - Produk",
		"Nama":       session["nama"],
		"data":       file,
		"barang":     barang,
		"namaBarang": r.URL.Query().Get("barang"),
	})
}

func (controller *ProdukControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("produk/views/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	barang, _ := controller.service.GetBahan(context.Background())
	myTemplate.ExecuteTemplate(w, "createProduk", map[string]interface{}{
		"Title":  "Cafe Mewa - Produk",
		"Nama":   session["nama"],
		"barang": barang,
	})
}

func (controller *ProdukControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	idUser := 1
	idBahan, err := strconv.Atoi(r.PostFormValue("barang"))
	helperMain.PanicIfError(err)
	harga, err := strconv.Atoi(r.PostFormValue("harga"))
	helperMain.PanicIfError(err)
	var stock bool
	if r.PostFormValue("stock") != "" {
		stock = true
	} else {
		stock = false
	}
	request := entity.Produk{
		Id_User:  idUser,
		Id_Bahan: idBahan,
		Harga:    harga,
		Stock:    stock,
	}
	check, err := controller.service.CheckByBahan(context.Background(), idBahan)
	if check.Id == 0 {
		controller.service.Create(context.Background(), request)
		helperMain.PanicIfError(err)
		http.Redirect(w, r, "/produk/", http.StatusFound)
	} else {
		controller.service.Update(context.Background(), request)
		http.Redirect(w, r, "/produk/", http.StatusFound)
	}
	return
}

func (controller *ProdukControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	idBahan, err := strconv.Atoi(r.PostFormValue("barang"))
	harga, err := strconv.Atoi(r.PostFormValue("harga"))
	helperMain.PanicIfError(err)
	var stock bool
	if r.PostFormValue("stock") != "" {
		stock = true
	} else {
		stock = false
	}
	request := entity.Produk{
		Id:       id,
		Id_User:  1,
		Id_Bahan: idBahan,
		Harga:    harga,
		Stock:    stock,
	}
	controller.service.Update(context.Background(), request)
	http.Redirect(w, r, "/produk/", http.StatusFound)
}

func (controller *ProdukControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	err = controller.service.Delete(context.Background(), id)
	helperMain.PanicIfError(err)
	http.Redirect(w, r, "/produk/", http.StatusAccepted)
}
