package web

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
