package controller

import (
	"context"
	"eh_teh_mewa/pembelian/helper"
	"eh_teh_mewa/pembelian/model/web"
	"eh_teh_mewa/pembelian/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type PembelianControllerImpl struct {
	PembelianService service.PembelianService
}

type Data struct {
	Title   string
	AllData []web.PembelianResponse
}

func NewPembelianController(pembelianService service.PembelianService) PembelianController {
	return &PembelianControllerImpl{
		PembelianService: pembelianService,
	}
}

func (controller *PembelianControllerImpl) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	myTemplate := template.New("LIST").Funcs(map[string]interface{}{
		"add":          helper.Add,
		"format":       helper.FormatTanggal,
		"formatNumber": helper.FormatNumberic,
	})
	myTemplate = template.Must(myTemplate.ParseFiles(
		"pembelian/views/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	awal := request.URL.Query().Get("awal")
	akhir := request.URL.Query().Get("akhir")
	serv, err := controller.PembelianService.FindByAll(context.Background(), awal, akhir)
	helper.PanicIfError(err)
	myTemplate.ExecuteTemplate(writer, "indexPembelian", Data{
		Title:   "Cafe Mewa",
		AllData: serv,
	})
}

func (controller *PembelianControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	myTemplate := template.New("INSERT")
	myTemplate = template.Must(myTemplate.ParseFiles(
		"pembelian/views/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(writer, "createPembelian", map[string]interface{}{
		"Title": "Cafe Mewa - Create",
		"barang": map[int]string{
			1: "kopi",
		},
	})
}
func (controller *PembelianControllerImpl) Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	barang, err := strconv.Atoi(request.PostFormValue("barang"))
	helper.PanicIfError(err)
	jumlah, err := strconv.Atoi(request.PostFormValue("jumlah"))
	helper.PanicIfError(err)
	biaya, err := strconv.Atoi(request.PostFormValue("biaya"))
	helper.PanicIfError(err)
	use := true
	if request.PostFormValue("use") == "on" {
		use = true
	} else {
		use = false
	}

	requestCustom := web.PembelianCreateRequest{
		Id_user:       1,
		Id_bahan_baku: barang,
		Jumlah:        jumlah,
		Tanggal:       request.PostFormValue("tanggal"),
		Biaya:         biaya,
		Use_pembelian: use,
	}
	respon, err := controller.PembelianService.Store(context.Background(), requestCustom)
	helper.PanicIfError(err)
	http.Redirect(writer, request, "/pembelian/show/"+strconv.Itoa(respon.Id), http.StatusFound)
	return
}
func (controller *PembelianControllerImpl) Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	myTemplate := template.New("UPDATE").Funcs(
		map[string]interface{}{
			"add":          helper.Add,
			"format":       helper.FormatTanggal,
			"formatNumber": helper.FormatNumberic,
		})
	myTemplate = template.Must(myTemplate.ParseFiles(
		"pembelian/views/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	serv, err := controller.PembelianService.FindById(context.Background(), params.ByName("id"))
	helper.PanicIfError(err)
	//fmt.Fprintln(writer, serv)
	type barang struct {
		id   int
		nama string
	}
	myTemplate.ExecuteTemplate(writer, "showPembelian", map[string]interface{}{
		"Title":  "Cafe Mewa - Create",
		"data":   serv,
		"barang": barang{1, "kopi"},
	})
}
func (controller *PembelianControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("Controller")
}
func (controller *PembelianControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := controller.PembelianService.Delete(context.Background(), params.ByName("id"))
	helper.PanicIfError(err)
	http.Redirect(writer, request, "/pembelian/", http.StatusFound)
	return
}
