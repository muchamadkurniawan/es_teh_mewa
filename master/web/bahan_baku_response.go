package web

type BahanbakuResponse struct {
	Id       int
	IdSatuan SatuanResponse
	Nama     string
}

type BahanbakuFullResponse struct {
	Id       int
	IdSatuan string
	Nama     string
}
