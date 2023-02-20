package controller

import (
	"context"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/produk/model/entity"
	"eh_teh_mewa/produk/model/web"
	"eh_teh_mewa/produk/service"
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

func (controller *ProdukControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	file, err := controller.service.FindById(context.Background(), id)
	helperMain.PanicIfError(err)
	myTemplate := template.Must(template.ParseFiles("produk/views/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "showProduk", map[string]interface{}{
		"Title": "Cafe Mewa - Produk",
		"data":  file,
	})
}

func (controller *ProdukControllerImpl) FindByAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var file []web.ResponseProdukFull
	var err error
	if params.ByName("barang") != "" {
		file, err = controller.service.FindByAll(context.Background())
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
		"data":       file,
		"barang":     barang,
		"namaBarang": r.URL.Query().Get("barang"),
	})
}

func (controller *ProdukControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("produk/views/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "indexProduk", map[string]interface{}{
		"Title": "Cafe Mewa - Produk",
	})
}

func (controller *ProdukControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idUser := 1
	idBahan, err := strconv.Atoi(r.PostFormValue("bahan"))
	harga, err := strconv.Atoi(r.PostFormValue("harga"))
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
	file, err := controller.service.Create(context.Background(), request)
	helperMain.PanicIfError(err)
	id := strconv.Itoa(file.Id)
	http.Redirect(w, r, "/produk/show/"+id+"/", http.StatusAccepted)
}

func (controller *ProdukControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	idBahan, err := strconv.Atoi(r.PostFormValue("bahan"))
	harga, err := strconv.Atoi(r.PostFormValue("harga"))
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
	_, err = controller.service.Update(context.Background(), request)
	helperMain.PanicIfError(err)
	http.Redirect(w, r, "/produk/show/"+params.ByName("id")+"/", http.StatusAccepted)
}

func (controller *ProdukControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	err = controller.service.Delete(context.Background(), id)
	helperMain.PanicIfError(err)
	http.Redirect(w, r, "/produk/", http.StatusAccepted)
}
