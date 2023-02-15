package web

type RequestProduk struct {
	Id_User   int  `json:"id_user"`
	Id_Barang int  `json:"id_barang"`
	Harga     int  `json:"harga"`
	Stock     bool `json:"stock"`
}
