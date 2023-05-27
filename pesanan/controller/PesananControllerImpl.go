package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/pesanan/service"
	"es_teh_mewa/pesanan/web"
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

func (controller *PesananControllerImpl) CheckLoginKasir(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "kasir" {
		http.Redirect(w, r, "/user/", http.StatusFound)
	}
	return session.Values
}

func (controller *PesananControllerImpl) CheckLoginAdmin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *PesananControllerImpl) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
	myTemplate := template.Must(template.ParseFiles("pesanan/views/show.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	detail := controller.service.FindPesananDetail(context.Background(), id)
	produks := controller.service.GetProdukJualsAll(context.Background())
	helperMain.PanicIfError(err)
	myTemplate.ExecuteTemplate(w, "showPesanan", map[string]interface{}{
		"Title":   "Cafe Mewa - Detail",
		"Nama":    session["nama"],
		"data":    detail,
		"produks": produks,
	})
}

func (controller *PesananControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
	myTemplate := template.Must(template.ParseFiles("pesanan/views/index.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml"))
	data := controller.service.FindAllPesananDetail(context.Background())
	myTemplate.ExecuteTemplate(w, "indexPesanan", map[string]interface{}{
		"Title": "Cafe Mewa - Order",
		"Nama":  session["nama"],
		"now":   time.Now().Format("02 Jan 2006"),
		"data":  data,
	})
}

func (controller *PesananControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
	myTemplate := template.Must(template.ParseFiles("pesanan/views/create.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
	))
	produk := controller.service.GetProdukJualsAll(context.Background())
	myTemplate.ExecuteTemplate(w, "createPesanan", map[string]interface{}{
		"Title":  "Cafe Mewa - List Order",
		"Nama":   session["nama"],
		"now":    time.Now().Format("02 Jan 2006"),
		"produk": produk,
	})
}

func (controller *PesananControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLoginKasir(w, r)
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
	var dataDetail []web.BiayaPesananReq
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
			err, idDetail := controller.service.CreateDetail(context.Background(), request)
			dataDetail = append(dataDetail, web.BiayaPesananReq{
				Id_produk: idProduk,
				Id_detail: idDetail,
				Jumlah:    jumlah,
			})
		}
	}
	for _, req := range dataDetail {
		bahan := controller.service.GetIdBahan(context.Background(), req.Id_produk)
		harga := controller.service.GetHargaBiaya(context.Background(), bahan)
		err := controller.service.CreateBiaya(context.Background(), bahan, req.Jumlah, harga, req.Id_detail)
		controller.service.UpdateStok(context.Background(), bahan, req.Jumlah)
		helperMain.PanicIfError(err)
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
	controller.CheckLoginKasir(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	err = controller.service.Cetak(context.Background(), id)
	if err != nil {
		fmt.Fprintln(w, "Error", err)
	} else {
		http.Redirect(w, r, "/storage/struk.pdf", http.StatusFound)
	}
}

func (controller *PesananControllerImpl) AddBiaya(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginKasir(w, r)
	myTemplate := template.Must(template.ParseFiles("pesanan/views/create.gohtml", "view/layout/app.gohtml",
		"view/kasir/headKasir.gohtml", "view/kasir/footerKasir.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
	))
	produk := controller.service.GetProdukJualsAll(context.Background())
	myTemplate.ExecuteTemplate(w, "createPesanan", map[string]interface{}{
		"Title":  "Cafe Mewa - List Order",
		"Nama":   session["nama"],
		"now":    time.Now().Format("02 Jan 2006"),
		"produk": produk,
	})
}

func (controller *PesananControllerImpl) ShowAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := controller.CheckLoginAdmin(w, r)
	myTemplate := template.Must(template.ParseFiles("pesanan/views/showAdmin.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	detail := controller.service.FindPesananDetail(context.Background(), id)
	produks := controller.service.GetProdukJualsAll(context.Background())
	helperMain.PanicIfError(err)
	myTemplate.ExecuteTemplate(w, "showPesananAdmin", map[string]interface{}{
		"Title":   "Cafe Mewa - Detail",
		"Nama":    session["nama"],
		"data":    detail,
		"produks": produks,
	})
}
