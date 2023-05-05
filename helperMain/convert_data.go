package helperMain

import (
	"encoding/json"
	"es_teh_mewa/master/web"
)

func StructToMap_Users(data web.UsersResponse) map[string]interface{} {
	var myMap map[string]interface{}
	dataConv, _ := json.Marshal(data)
	json.Unmarshal(dataConv, &myMap)
	return myMap
}

func StructSliceToMap_Users(datas []web.UsersResponse) []map[string]interface{} {
	var dataMAPs []map[string]interface{}
	for index, _ := range datas {
		data := datas[index]
		var myMap map[string]interface{}
		dataConv, _ := json.Marshal(data)
		json.Unmarshal(dataConv, &myMap)
		dataMAPs = append(dataMAPs, myMap)
	}
	return dataMAPs
}

func StructToMap_Satuan(data web.SatuanResponse) map[string]interface{} {
	var myMap map[string]interface{}
	dataConv, _ := json.Marshal(data)
	json.Unmarshal(dataConv, &myMap)
	return myMap
}

func StructSliceToMap_Satuan(datas []web.SatuanResponse) []map[string]interface{} {
	var dataMAPs []map[string]interface{}
	for index, _ := range datas {
		data := datas[index]
		var myMap map[string]interface{}
		dataConv, _ := json.Marshal(data)
		json.Unmarshal(dataConv, &myMap)
		dataMAPs = append(dataMAPs, myMap)
	}
	return dataMAPs
}

func StructToMap_Bahan(data web.BahanbakuResponse) map[string]interface{} {
	var myMap map[string]interface{}
	dataConv, _ := json.Marshal(data)
	json.Unmarshal(dataConv, &myMap)
	return myMap
}

func StructSliceToMap_Bahan(datas []web.BahanbakuResponse) []map[string]interface{} {
	var dataMAPs []map[string]interface{}
	for index, _ := range datas {
		data := datas[index]
		var myMap map[string]interface{}
		dataConv, _ := json.Marshal(data)
		json.Unmarshal(dataConv, &myMap)
		dataMAPs = append(dataMAPs, myMap)
	}
	return dataMAPs
}
