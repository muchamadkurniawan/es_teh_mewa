package web

import "time"

type ResponseDashboard struct {
	Id         int
	Tanggal    time.Time
	Keterangan string
}
