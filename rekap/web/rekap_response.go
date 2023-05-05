package web

import "time"

type RekapResponseTime struct {
	id         int
	tanggal    time.Time
	keterangan string
}

type RekapResponseDateString struct {
	id         int
	tanggal    string
	keterangan string
}

type AllPesananRekap struct {
	Id           int
	Tanggal      time.Time
	Pembayaran   bool
	TotalPesanan int
}
