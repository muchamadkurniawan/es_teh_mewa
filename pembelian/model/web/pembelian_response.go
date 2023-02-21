package web

import "time"

type PembelianResponse struct {
	Id            int
	Id_user       int
	Id_bahan_baku int
	Tanggal       string
	Biaya         int
	Jumlah        int
	Use_pembelian bool
}

type PembelianResponseFull struct {
	Id            int
	Id_user       int
	Id_bahan_baku string
	Tanggal       time.Time
	Biaya         int
	Jumlah        int
	Use_pembelian bool
}
