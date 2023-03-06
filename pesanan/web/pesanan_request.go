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
