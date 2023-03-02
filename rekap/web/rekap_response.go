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
