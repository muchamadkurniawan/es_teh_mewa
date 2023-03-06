package helper

import (
	"eh_teh_mewa/pesanan/model/entity"
	"eh_teh_mewa/pesanan/web"
)

func PesananEntityToResponseString(pesanan entity.PesananEntity) web.PesananResponseDateString {
	tgl := pesanan.Tanggal.Format("02 Jan 2006")
	respon := web.PesananResponseDateString{
		Id:         pesanan.Id,
		Id_user:    pesanan.Id_user,
		Id_rekap:   pesanan.Id_rekap,
		Tanggal:    tgl,
		Pembayaran: pesanan.Pembayaran,
	}
	return respon
}
