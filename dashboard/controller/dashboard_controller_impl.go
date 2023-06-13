package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/dashboard/service"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/pesanan/helper"
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type DashboardControllerImpl struct {
	service service.DashboardService
}

func NewDashboardController(service service.DashboardService) DashboardContrller {
	return &DashboardControllerImpl{
		service: service,
	}
}

func (controller *DashboardControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (d DashboardControllerImpl) GetRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := d.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"dashboard/view/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	data := d.service.GetRekap(context.Background())
	myTemplate.ExecuteTemplate(w, "indexDashboard", map[string]interface{}{
		"Title": "Cafe Mewa",
		"Nama":  session["nama"],
		"Data":  data,
	})
}

func (d *DashboardControllerImpl) DetailRekap(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session := d.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"dashboard/view/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	data := d.service.GetRekapById(context.Background(), id)
	pesanan, totalPesanan := d.service.GetPesananRekapById(context.Background(), id)
	biaya, totalBiaya := d.service.GetBiayaRekapById(context.Background(), id)
	myTemplate.ExecuteTemplate(w, "showDashboard", map[string]interface{}{
		"Title":        "Cafe Mewa",
		"Nama":         session["nama"],
		"Data":         data,
		"Pesanan":      pesanan,
		"Biaya":        biaya,
		"TotalPesanan": totalPesanan,
		"TotalBiaya":   totalBiaya,
		"Laba":         totalPesanan - totalBiaya,
	})
}

func (controller *DashboardControllerImpl) Cetak(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	m := pdf.NewMaroto(consts.Portrait, consts.A5)
	m.SetPageMargins(20, 10, 20)
	helper.PdfHeader(m)
	tableHeading := []string{"Tanggal", "Kasir"}
	m.SetBackgroundColor(color.NewWhite())
	data := controller.service.GetRekapById(context.Background(), id)
	converData := helper.SetData2(data)
	m.Row(5.0, func() {
		m.Text("Rekapulasi", props.Text{VerticalPadding: 10.0})
	})
	m.Row(20.0, func() {
		m.TableList(tableHeading, converData, props.TableList{
			HeaderProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 4},
			},
			ContentProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 4},
			},

			Align:              consts.Left,
			HeaderContentSpace: 2,
			Line:               true,
		})
	})
	m.Row(6.0, func() {

	})
	m.Row(5.0, func() {
		m.Text("Pesanan")
	})
	tableHeading2 := []string{"Tanggal", "Total", "Pembayaran"}
	data2, total2 := controller.service.GetPesananRekapById(context.Background(), id)
	converData2 := helper.SetData4(data2, total2)
	m.Row(10.0, func() {
		m.TableList(tableHeading2, converData2, props.TableList{
			HeaderProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 3, 3},
			},
			ContentProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 3, 3},
			},

			Align:              consts.Left,
			HeaderContentSpace: 2,
			Line:               true,
		})
	})
	m.Row(6.0, func() {

	})
	m.Row(5.0, func() {
		m.Text("Biaya")
	})
	tableHeading3 := []string{"Barang", "Jumlah", "Harga", "Total"}
	data3, total3 := controller.service.GetBiayaRekapById(context.Background(), id)
	converData3 := helper.SetData3(data3, total3)
	m.Row(10.0, func() {
		m.TableList(tableHeading3, converData3, props.TableList{
			HeaderProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 2, 2, 2},
			},
			ContentProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 2, 2, 2},
			},

			Align:              consts.Left,
			HeaderContentSpace: 2,
			Line:               true,
		})
	})

	m.Row(6.0, func() {

	})
	m.Row(1.0, func() {
		m.Text("Laba Rugi")
	})
	tableHeading4 := []string{"		", "	", "	", "	"}
	var datas [][]string
	laba := []string{"Total	", "	", strconv.Itoa(total2 - total3)}
	datas = append(datas, laba)
	m.Row(10.0, func() {
		m.TableList(tableHeading4, datas, props.TableList{
			HeaderProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 2, 2, 2},
			},
			ContentProp: props.TableListContent{
				Size:      9,
				GridSizes: []uint{4, 2, 2, 2},
			},

			Align:              consts.Left,
			HeaderContentSpace: 2,
			Line:               true,
		})
	})
	err = m.OutputFileAndClose("storage/rekap.pdf")
	//err = controller.service.Cetak(context.Background(), id)
	if err != nil {
		fmt.Fprintln(w, "Error", err)
	} else {
		http.Redirect(w, r, "/storage/rekap.pdf", http.StatusFound)
	}
}
