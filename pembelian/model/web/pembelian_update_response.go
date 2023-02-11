package web

import "time"

type PembelianUpdateResponse struct {
	Id            int
	Id_user       int
	Id_bahan_baku int
	Tanggal       time.Time
	Biaya         int
	Jumlah        int
	Use_pembelian bool
}
