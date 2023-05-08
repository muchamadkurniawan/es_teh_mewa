package web

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
