package helper

import (
	web3 "es_teh_mewa/biaya/web"
	web2 "es_teh_mewa/dashboard/web"
	"es_teh_mewa/pesanan/web"
	web4 "es_teh_mewa/rekap/web"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"strconv"
)

func PdfBlankColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
}

func PdfHeader(m pdf.Maroto) {
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Cafe Mawah",
				props.Text{
					Top:   3,
					Style: consts.Bold,
					Align: consts.Left,
					Color: PdfBlankColor(),
				})
		})
	})
}

func SetData(data []web.DetailRespon) [][]string {
	var datas [][]string
	total := 0
	for _, datum := range data {
		a := []string{}
		a = append(a, datum.Nama_produk)
		a = append(a, strconv.Itoa(datum.Jumlah))
		a = append(a, strconv.Itoa(datum.Harga))
		a = append(a, strconv.Itoa(datum.Total))
		datas = append(datas, a)
		total += datum.Total
	}
	datas = append(datas, []string{"Total	", "	", "	", strconv.Itoa(total)})
	return datas
}

func SetData2(data web2.ResponseDashboard) [][]string {
	var datas [][]string
	a := []string{}
	a = append(a, data.Tanggal.Format("02 January 2006"))
	a = append(a, data.Keterangan)
	datas = append(datas, a)
	return datas
}

func SetData3(data []web3.GetBiayaTodayRespon, total int) [][]string {
	var datas [][]string
	for _, datum := range data {
		a := []string{}
		a = append(a, datum.Barang)
		a = append(a, strconv.Itoa(datum.Jumlah))
		a = append(a, strconv.Itoa(datum.HargaSatuan))
		a = append(a, strconv.Itoa(datum.Total))
		datas = append(datas, a)
	}
	datas = append(datas, []string{"Total	", "	", "	", strconv.Itoa(total)})
	return datas
}

func SetData4(data []web4.AllPesananRekap, total int) [][]string {
	var datas [][]string
	var bayar string
	for _, datum := range data {
		if datum.Pembayaran {
			bayar = "Lunas"
		} else {
			bayar = "Belum"
		}
		a := []string{}
		a = append(a, datum.Tanggal.Format("02 January 2006"))
		a = append(a, strconv.Itoa(datum.TotalPesanan))
		a = append(a, bayar)
		datas = append(datas, a)
	}
	datas = append(datas, []string{"Total	", "	", strconv.Itoa(total)})
	return datas
}
