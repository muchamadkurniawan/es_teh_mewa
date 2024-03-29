package web

import "time"

type PesananRequestDateString struct {
	Id_user    int
	Id_rekap   int
	Tanggal    string
	Pembayaran bool
}

type PesananRequest struct {
	Id_user    int
	Id_rekap   int
	Tanggal    time.Time
	Pembayaran bool
}

type PesananNamaBarang struct {
	Id_detail    int
	Id_pesanan   int
	Nama         string
	Harga_satuan int
	Harga        int
}

type PesananRequestUpdate struct {
	Id         int
	Date       time.Time
	Pembayaran bool
}

type PesananRequestSum struct {
	Id         int
	Time       string
	Pembayaran bool
	Total      int
}

type PesananDetail struct {
	Pesanan PesananRequestSum
	Detail  []PesananNamaBarang
}

type PesananDetailUpdate struct {
	Pesanan PesananRequestUpdate
	Detail  []DetailRespon
}

type CetakPesanan struct {
	Keterangan string
	Harga      int
	Total      int
}

type BiayaPesananReq struct {
	Id_produk int
	Id_detail int
	Jumlah    int
}
