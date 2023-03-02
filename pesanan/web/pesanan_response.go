package web

import "time"

type PesananResponseDateString struct {
	id       int
	id_user  int
	id_rekap int
	tanggal  string
}

type PesananResponse struct {
	id       int
	id_user  int
	id_rekap int
	tanggal  time.Time
}
