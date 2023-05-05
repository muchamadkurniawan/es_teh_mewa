package helper

import (
	"es_teh_mewa/helperMain"
	"es_teh_mewa/pembelian/model/entity"
	"es_teh_mewa/pembelian/model/web"
	"time"
)

func StringToTime(date string) time.Time {
	dateString := date + "00:00"
	myDate, err := time.Parse("2006-01-02 15:04", dateString)
	helperMain.PanicIfError(err)
	return myDate
}

func RequestToResponse(pembelian web.PembelianCreateRequest) web.PembelianResponse {
	return web.PembelianResponse{
		Id:            int(pembelian.Id),
		Id_user:       int(pembelian.Id_user),
		Id_bahan_baku: int(pembelian.Id_bahan_baku),
		Tanggal:       pembelian.Tanggal,
		Biaya:         int(pembelian.Biaya),
		Jumlah:        int(pembelian.Jumlah),
		Use_pembelian: pembelian.Use_pembelian,
	}
}

func ToPembelianResponse(pembelian entity.Pembelian) web.PembelianResponse {
	return web.PembelianResponse{
		Id:            int(pembelian.Id),
		Id_user:       int(pembelian.IdUser),
		Id_bahan_baku: int(pembelian.IdBahan_Baku),
		Tanggal:       FormatTanggal(pembelian.Tanggal),
		Biaya:         int(pembelian.Biaya),
		Jumlah:        int(pembelian.Jumlah),
		Use_pembelian: pembelian.UsePembelian,
	}
}

func ToPembelianUpdateResponse(pembelian entity.Pembelian) web.PembelianUpdateResponse {
	return web.PembelianUpdateResponse{
		Id:            int(pembelian.Id),
		Id_user:       int(pembelian.IdUser),
		Id_bahan_baku: int(pembelian.IdBahan_Baku),
		Tanggal:       pembelian.Tanggal,
		Biaya:         int(pembelian.Biaya),
		Jumlah:        int(pembelian.Jumlah),
		Use_pembelian: pembelian.UsePembelian,
	}
}

func ToEntityPembelian(pembelian web.PembelianCreateRequest) entity.Pembelian {
	return entity.Pembelian{
		IdUser:       int32(pembelian.Id_user),
		IdBahan_Baku: int32(pembelian.Id_bahan_baku),
		Tanggal:      StringToTime(pembelian.Tanggal),
		Biaya:        int32(pembelian.Biaya),
		Jumlah:       int32(pembelian.Jumlah),
		UsePembelian: pembelian.Use_pembelian,
	}
}

//func ToPembelianUpdateRequest(pembelian web.PembelianUpdateRequest) entity.Pembelian {
//	return entity.Pembelian{
//		Id:           int32(pembelian.Id),
//		IdUser:       int32(pembelian.Id_user),
//		IdBahan_Baku: int32(pembelian.Id_bahan_baku),
//		Tanggal:      pembelian.Tanggal,
//		Biaya:        int32(pembelian.Biaya),
//		Jumlah:       int32(pembelian.Jumlah),
//		UsePembelian: pembelian.Use_pembelian,
//	}
//}

func ToPembelianAllResponse(pembelian []entity.Pembelian) []web.PembelianResponse {
	var pembelianResponse []web.PembelianResponse
	for _, v := range pembelian {
		pembelianResponse = append(pembelianResponse, ToPembelianResponse(v))
	}
	return pembelianResponse
}

//func ToPembelianRequest(pembelian entity.Pembelian) web.PembelianCreateRequest {
//	return web.PembelianCreateRequest{
//		Id_user:       int(pembelian.IdUser),
//		Id_bahan_baku: int(pembelian.IdBahan_Baku),
//		Tanggal:       pembelian.Tanggal,
//		Biaya:         int(pembelian.Biaya),
//		Jumlah:        int(pembelian.Jumlah),
//		Use_pembelian: pembelian.UsePembelian,
//	}
//}
