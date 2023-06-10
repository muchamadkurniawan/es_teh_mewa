package entity

import "time"

type Pembelian struct {
	Id, IdUser, IdBahan_Baku, Biaya, Jumlah, Total int32
	Tanggal                                        time.Time
	UsePembelian                                   bool
}
