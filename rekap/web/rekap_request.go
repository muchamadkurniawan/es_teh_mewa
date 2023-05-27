package web

type RekapRequestDateString struct {
	Tanggal    string
	Keterangan string
}

type DetailPesanan struct {
	Id        int
	Id_Produk int
	Jumlah    int
	Harga     int
	Total     int
}
