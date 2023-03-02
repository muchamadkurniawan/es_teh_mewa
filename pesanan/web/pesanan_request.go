package web

import "time"

type PesananRequestDateString struct {
	id_user  int
	id_rekap int
	tanggal  string
}

type PesananRequest struct {
	id_user  int
	id_rekap int
	tanggal  time.Time
}
