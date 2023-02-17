package helperMain

import (
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/web"
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
		Type:     user.Type_user,
	}
}

func ToSatuanResponse(satuan entity.Satuan) web.SatuanResponse {
	return web.SatuanResponse{
		Id:   (satuan.Id),
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
func ToBahanRresponses(bahans []entity.BahanBaku) []web.BahanbakuResponse {
	var bahanResponses []web.BahanbakuResponse
	for _, bahan := range bahans {
		bahanResponses = append(bahanResponses, ToBahanResponse(bahan))
	}
	return bahanResponses
}
func ToBahanResponse(bahan entity.BahanBaku) web.BahanbakuResponse {
	return web.BahanbakuResponse{
		Id:       bahan.Id,
		IdSatuan: bahan.IdSatuan,
		Nama:     bahan.Nama,
	}
}
