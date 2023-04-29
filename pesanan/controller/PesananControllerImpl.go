package controller

import (
	"context"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/pesanan/service"
	"eh_teh_mewa/pesanan/web"
	"fmt"
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
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	detail := controller.service.FindPesananDetail(context.Background(), id)
	produks := controller.service.GetProdukJualsAll(context.Background())
	helperMain.PanicIfError(err)
	myTemplate.ExecuteTemplate(w, "showPesanan", map[string]interface{}{
		"Title":   "Cafe Mewa - Detail",
		"data":    detail,
		"produks": produks,
	})
}

func (controller *PesananControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("pesanan/views/index.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	data := controller.service.FindAllPesananDetail(context.Background())
	myTemplate.ExecuteTemplate(w, "indexPesanan", map[string]interface{}{
		"Title": "Cafe Mewa - Order",
		"now":   time.Now().Format("02 Jan 2006"),
		"data":  data,
	})
}

func (controller *PesananControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	myTemplate := template.Must(template.ParseFiles("pesanan/views/create.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
	))
	produk := controller.service.GetProdukJualsAll(context.Background())
	myTemplate.ExecuteTemplate(w, "createPesanan", map[string]interface{}{
		"Title":  "Cafe Mewa - List Order",
		"now":    time.Now().Format("02 Jan 2006"),
		"produk": produk,
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
		Pembayaran: bayar,
	}
	idPesanan := controller.service.CreatePesanan(context.Background(), data)
	ids := controller.service.GetIdProduk(context.Background())
	check := true
	for _, id := range ids {
		if r.PostFormValue("qty"+id) != "" {
			check = false
			idProduk, err := strconv.Atoi(r.PostFormValue("barang" + id))
			harga, err := strconv.Atoi(r.PostFormValue("harga" + id))
			jumlah, err := strconv.Atoi(r.PostFormValue("qty" + id))
			helperMain.PanicIfError(err)
			total := harga * jumlah
			request := web.DetailRequest{
				Id_produk:  idProduk,
				Id_pesanan: idPesanan,
				Harga:      harga,
				Jumlah:     jumlah,
				Total:      total,
			}
			err = controller.service.CreateDetail(context.Background(), request)
		}
	}
	if check {
		http.Redirect(w, r, "/pesanan/order/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
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

func (controller *PesananControllerImpl) Cetak(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	err = controller.service.Cetak(context.Background(), id)
	if err != nil {
		fmt.Fprintln(w, "Error", err)
	} else {
		http.Redirect(w, r, "/storage/struk.pdf", http.StatusFound)
	}
}
