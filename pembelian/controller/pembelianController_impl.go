package controller

import (
	"context"
	"eh_teh_mewa/auth/config"
	"eh_teh_mewa/helperMain"
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

func (controller *PembelianControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *PembelianControllerImpl) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(writer, request)
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
	helperMain.PanicIfError(err)
	myTemplate.ExecuteTemplate(writer, "indexPembelian", map[string]interface{}{
		"Title":   "Cafe Mewa",
		"Nama":    session["nama"],
		"Awal":    awal,
		"Akhir":   akhir,
		"AllData": serv,
	})
}

func (controller *PembelianControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(writer, request)
	myTemplate := template.New("INSERT")
	myTemplate = template.Must(myTemplate.ParseFiles(
		"pembelian/views/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	bahan := controller.PembelianService.GetAllBahanBaku(context.Background())
	myTemplate.ExecuteTemplate(writer, "createPembelian", map[string]interface{}{
		"Title":  "Cafe Mewa - Create",
		"Nama":   session["nama"],
		"barang": bahan,
	})
}
func (controller *PembelianControllerImpl) Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_ = controller.CheckLogin(writer, request)
	barang, err := strconv.Atoi(request.PostFormValue("barang"))
	helperMain.PanicIfError(err)
	jumlah, err := strconv.Atoi(request.PostFormValue("jumlah"))
	helperMain.PanicIfError(err)
	biaya, err := strconv.Atoi(request.PostFormValue("biaya"))
	helperMain.PanicIfError(err)
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
	helperMain.PanicIfError(err)
	http.Redirect(writer, request, "/pembelian/show/"+strconv.Itoa(respon.Id), http.StatusFound)
	return
}
func (controller *PembelianControllerImpl) Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	session := controller.CheckLogin(writer, request)
	myTemplate := template.New("UPDATE").Funcs(
		map[string]interface{}{
			"add":          helper.Add,
			"format":       helper.FormatTanggalUpdate,
			"formatNumber": helper.FormatNumberic,
			"checkbool":    helper.CheckBool,
		})
	myTemplate = template.Must(myTemplate.ParseFiles(
		"pembelian/views/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	serv, err := controller.PembelianService.FindById(context.Background(), params.ByName("id"))
	helperMain.PanicIfError(err)
	bahan := controller.PembelianService.GetAllBahanBaku(context.Background())
	myTemplate.ExecuteTemplate(writer, "showPembelian", map[string]interface{}{
		"Title":  "Cafe Mewa - Create",
		"Nama":   session["nama"],
		"data":   serv,
		"barang": bahan,
	})
}
func (controller *PembelianControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_ = controller.CheckLogin(writer, request)
	id, err := strconv.Atoi(params.ByName("id"))
	barang, err := strconv.Atoi(request.PostFormValue("barang"))
	jumlah, err := strconv.Atoi(request.PostFormValue("jumlah"))
	biaya, err := strconv.Atoi(request.PostFormValue("biaya"))
	use := true
	if request.PostFormValue("use") == "on" {
		use = true
	} else {
		use = false
	}
	helperMain.PanicIfError(err)
	respon, err := controller.PembelianService.Update(context.Background(), web.PembelianCreateRequest{
		Id:            id,
		Id_user:       1,
		Id_bahan_baku: barang,
		Tanggal:       request.PostFormValue("tanggal"),
		Jumlah:        jumlah,
		Biaya:         biaya,
		Use_pembelian: use,
	})
	helperMain.PanicIfError(err)
	http.Redirect(writer, request, "/pembelian/show/"+strconv.Itoa(respon.Id), http.StatusFound)
	return
}
func (controller *PembelianControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_ = controller.CheckLogin(writer, request)
	err := controller.PembelianService.Delete(context.Background(), params.ByName("id"))
	helperMain.PanicIfError(err)
	http.Redirect(writer, request, "/pembelian/", http.StatusFound)
	return
}
