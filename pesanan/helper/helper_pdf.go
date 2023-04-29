package helper

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
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
					Align: consts.Center,
					Color: PdfBlankColor(),
				})
		})
	})
}
