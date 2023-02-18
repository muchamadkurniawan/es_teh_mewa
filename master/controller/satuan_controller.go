package controller

import (
	"context"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/service"
	"eh_teh_mewa/master/web"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type SatuanControllerImpl struct {
	service service.SatuanService
}

func NewSatuanController(serviceSatuan service.SatuanService) SatuanController {
	return &SatuanControllerImpl{
		service: serviceSatuan,
	}
}

func (controller *SatuanControllerImpl) FindAllSatuan(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	serv := controller.service.FindAll(context.Background())
	myTemplate := template.Must(template.ParseFiles("master/views/satuan/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))

	myTemplate.ExecuteTemplate(response, "indexSatuan", map[string]interface{}{
		"data":  serv,
		"Title": "Cafe Mewa - Satuan",
	})
}

func (controller *SatuanControllerImpl) FindById(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	serv := controller.service.FindById(context.Background(), id)
	myTemplate := template.Must(template.ParseFiles("master/views/satuan/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))

	myTemplate.ExecuteTemplate(response, "showSatuan", map[string]interface{}{
		"data":  serv,
		"Title": "Cafe Mewa - Satuan",
	})
}

func (controller *SatuanControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles("master/views/satuan/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))

	myTemplate.ExecuteTemplate(w, "showSatuan", map[string]interface{}{
		"Title": "Cafe Mewa - Satuan",
	})
}

func (controller *SatuanControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	file := web.SatuanRequest{
		Nama: r.PostFormValue("name"),
	}
	controller.service.Save(context.Background(), file)
	http.Redirect(w, r, "/satuan/", http.StatusFound)
}

func (controller *SatuanControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	file := web.SatuanResponse{
		Id:   id,
		Name: r.PostFormValue("name"),
	}
	controller.service.Update(context.Background(), file)
	http.Redirect(w, r, "/satuan/", http.StatusFound)
}

func (controller *SatuanControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helperMain.PanicIfError(err)
	controller.service.Delete(context.Background(), id)
	http.Redirect(w, r, "/satuan/", http.StatusFound)
}
