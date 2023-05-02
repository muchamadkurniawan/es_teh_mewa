package helper

import (
	"eh_teh_mewa/pesanan/web"
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
		a = append(a, datum.Nama_produk+"	"+strconv.Itoa(datum.Jumlah))
		a = append(a, strconv.Itoa(datum.Harga))
		a = append(a, strconv.Itoa(datum.Total))
		datas = append(datas, a)
		total += datum.Total
	}
	datas = append(datas, []string{"Total	", "	", strconv.Itoa(total)})
	return datas
}
