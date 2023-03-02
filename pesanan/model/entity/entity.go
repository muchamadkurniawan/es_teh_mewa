package entity

import "time"

type PesananEntity struct {
	Id         int
	Id_user    int
	Id_rekap   int
	Tanggal    time.Time
	Pembayaran bool
}

type DetailPesananEntity struct {
	Id         int
	Id_produk  int
	Id_pesanan int
	Jumlah     int
	Harga      int
	Total      int
}
