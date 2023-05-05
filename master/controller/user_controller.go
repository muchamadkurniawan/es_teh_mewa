package controller

import (
	"context"
	"es_teh_mewa/auth/config"
	"es_teh_mewa/helperMain"
	service "es_teh_mewa/master/service"
	"es_teh_mewa/master/web"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UsersService
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &UserControllerImpl{UserService: usersService}
}

func (controller *UserControllerImpl) CheckLogin(w http.ResponseWriter, r *http.Request) map[interface{}]interface{} {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
	}
	if session.Values["type"] != "admin" {
		http.Redirect(w, r, "/pesanan/", http.StatusFound)
	}
	return session.Values
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles(
		"master/views/user/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "createUser", map[string]interface{}{
		"type": []string{
			"admin", "kasir",
		},
		"Nama":  session["nama"],
		"Title": "Cafe Mewa - User",
	})
}

func (controller *UserControllerImpl) Store(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//controller.CheckLogin(w, r)
	name := r.PostFormValue("username")
	password := r.PostFormValue("password")
	tipe := r.PostFormValue("typeUser")
	file := web.UsersCreateRequest{
		Username: name,
		Password: password,
		Type:     tipe,
	}
	controller.UserService.Create(context.Background(), file)
	http.Redirect(w, r, "/user/", http.StatusFound)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(param.ByName("id"))
	helperMain.PanicIfError(err)
	name := r.PostFormValue("username")
	tipe := r.PostFormValue("type")
	file := web.UsersResponse{
		Id:       id,
		Username: name,
		Tipe:     tipe,
	}
	controller.UserService.Update(context.Background(), file)
	http.Redirect(w, r, "/user/", http.StatusFound)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	controller.CheckLogin(w, r)
	id, err := strconv.Atoi(param.ByName("id"))
	helperMain.PanicIfError(err)
	controller.UserService.Delete(context.Background(), id)
	http.Redirect(w, r, "/user/", http.StatusFound)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("master/views/user/index.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))

	serv := controller.UserService.FindAll(context.Background())
	myTemplate.ExecuteTemplate(w, "indexUser", map[string]interface{}{
		"Title": "Cafe Mewa - User",
		"Nama":  session["nama"],
		"data":  serv,
	})
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	session := controller.CheckLogin(w, r)
	myTemplate := template.Must(template.ParseFiles("master/views/user/show.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	id, err := strconv.Atoi(param.ByName("id"))
	helperMain.PanicIfError(err)
	serv := controller.UserService.FindById(context.Background(), id)
	myTemplate.ExecuteTemplate(w, "showUser", map[string]interface{}{
		"data":  serv,
		"Title": "Cafe Mewa - User",
		"Nama":  session["nama"],
		"tipe": [2]string{
			"admin", "kasir",
		},
	})
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	myTemplate := template.Must(template.ParseFiles(
		"master/views/user/create.gohtml", "view/layout/app.gohtml",
		"view/layout/bodyTop.gohtml", "view/layout/footer.gohtml", "view/layout/head.gohtml", "view/layout/header.gohtml",
		"view/layout/sidebar.gohtml"))
	myTemplate.ExecuteTemplate(w, "createUser", map[string]interface{}{
		"type": []string{
			"admin", "kasir",
		},
		"Nama":  "Unknow",
		"Title": "Cafe Mewa - User",
	})
}
