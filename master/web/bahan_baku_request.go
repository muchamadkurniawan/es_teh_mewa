package web

type BahanbakuRequest struct {
	IdSatuan int    `json:"id_satuan"`
	Nama     string `json:"nama"`
}

type BahanbakuResponse struct {
	Id       int    `json:"id"`
	IdSatuan int    `json:"id_satuan"`
	Nama     string `json:"nama"`
}
