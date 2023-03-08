package controller

import (
	"context"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/pesanan/service"
	"eh_teh_mewa/pesanan/web"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type PesananControllerImpl struct {
	service service.PesananService
}

func NewPesananController(pesananService service.PesananService) PesananController {
	return &PesananControllerImpl{
		service: pesananService,
	}
}

func (controller *PesananControllerImpl) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("pesanan/views/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	data, err := controller.service.FindById(context.Background(), id)
	helperMain.PanicIfError(err)
	myTemplate.ExecuteTemplate(w, "showPesanan", map[string]interface{}{
		"Title": "Cafe Mewa - Detail",
		"data":  data,
	})
}

func (controller *PesananControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("pesanan/views/index.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
	))
	data, err := controller.service.FindAll(context.Background())
	produk := controller.service.GetProdukJualsAll(context.Background())
	helperMain.PanicIfError(err)
	//fmt.Fprintln(w, data, produk)
	myTemplate.ExecuteTemplate(w, "indexPesanan", map[string]interface{}{
		"Title":  "Cafe Mewa - List Order",
		"now":    time.Now().Format("02 Jan 2006"),
		"data":   data,
		"produk": produk,
	})
}

func (controller *PesananControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("pesanan/views/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "createPesanan", map[string]interface{}{
		"Title": "Cafe Mewa - Order",
	})
}

func (controller *PesananControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var bayar bool
	if r.PostFormValue("pembayaran") == "on" {
		bayar = true
	} else {
		bayar = false
	}
	data := web.PesananRequestDateString{
		Id_user:    1,
		Tanggal:    time.Now().Format("2006-01-02"),
		Pembayaran: bayar,
	}
	idPesanan := controller.service.CreatePesanan(context.Background(), data)
	ids := controller.service.GetIdProduk(context.Background())
	check := false
	for _, id := range ids {
		if r.PostFormValue("qty"+id) != "0" {
			check = true

		}
	}
}

func (controller *PesananControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *PesananControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
