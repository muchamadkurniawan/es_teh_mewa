package web

import "time"

type GetBahanBakuNonProdukRespon struct {
	Id     int
	Satuan string
	Nama   string
}

type GetBiayaTodayRespon struct {
	Id          int
	Barang      string
	Jumlah      int
	HargaSatuan int
	Total       int
}

type GetBiayaRespon struct {
	Id          int
	Barang      int
	Jumlah      int
	HargaSatuan int
	Total       int
	Id_Rekap    int
	Tanggal     time.Time
}
