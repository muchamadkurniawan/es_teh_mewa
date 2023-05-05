package entity

import "time"

type RekapEntity struct {
	id         int
	tanggal    time.Time
	keterangan string
}
