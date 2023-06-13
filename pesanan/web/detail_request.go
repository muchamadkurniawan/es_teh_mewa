package web

import "time"

type DetailRequest struct {
	Id_produk  int
	Id_pesanan int
	Jumlah     int
	Harga      int
	Total      int
}

type DetailRespon struct {
	Id          int
	Id_produk   int
	Nama_produk string
	Jumlah      int
	Harga       int
	Total       int
}

type DetailPesanan struct {
	Id        int
	Id_Produk int
	Jumlah    int
	Harga     int
	Total     int
}

type DetailRekap struct {
	Id         int
	Keterangan string
	Tanggal    time.Time
}
