package helperMain

import (
	"es_teh_mewa/master/model/entity"
	"es_teh_mewa/master/web"
)

func ToUserResponses(users []entity.Users) []web.UsersResponse {
	var userResponses []web.UsersResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:       int(user.Id),
		Username: user.UserName,
		Password: user.Password,
		Tipe:     user.Type_user,
	}
}

func ToSatuanResponse(satuan entity.Satuan) web.SatuanResponse {
	return web.SatuanResponse{
		Id:   int(satuan.Id),
		Name: satuan.Nama,
	}
}

func ToSatuanResponses(satuans []entity.Satuan) []web.SatuanResponse {
	var satuanResponses []web.SatuanResponse
	for _, satuan := range satuans {
		satuanResponses = append(satuanResponses, ToSatuanResponse(satuan))
	}
	return satuanResponses
}
func ToBahanRresponses(bahans []entity.BahanBakuFull) []web.BahanbakuFullResponse {
	var bahanResponses []web.BahanbakuFullResponse
	for _, bahan := range bahans {
		bahanResponses = append(bahanResponses, ToBahanrespon(bahan))
	}
	return bahanResponses
}

func ToBahanrespon(bahan entity.BahanBakuFull) web.BahanbakuFullResponse {
	newBahan := web.BahanbakuFullResponse{
		Id:       bahan.Id,
		IdSatuan: bahan.IdSatuan,
		Nama:     bahan.Nama,
	}
	return newBahan
}
