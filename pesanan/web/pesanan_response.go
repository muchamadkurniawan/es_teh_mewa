package web

import "time"

type PesananResponseDateString struct {
	Id         int
	Id_user    int
	Id_rekap   int
	Tanggal    string
	Pembayaran bool
}

type PesananResponse struct {
	Id         int
	Id_user    int
	Id_rekap   int
	Tanggal    time.Time
	Pembayaran bool
}
