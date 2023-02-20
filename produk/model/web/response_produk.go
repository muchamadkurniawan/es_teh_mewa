package web

type ResponseProduk struct {
	Id        int  `json:"id"`
	Id_User   int  `json:"id_user"`
	Id_Barang int  `json:"id_barang"`
	Harga     int  `json:"harga"`
	Stock     bool `json:"stock"`
}

type ResponseProdukFull struct {
	Id     int
	User   string
	Barang string
	Harga  int
	Stock  bool
}
