package controller

import (
	"context"
	"eh_teh_mewa/auth/config"
	"eh_teh_mewa/auth/model"
	"eh_teh_mewa/auth/service"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type LoginControllerImpl struct {
	LoginService service.LoginService
}

func NewLoginController(loginservice service.LoginService) LoginController {
	return &LoginControllerImpl{
		LoginService: loginservice,
	}
}

func (Controller *LoginControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method == http.MethodGet {
		myTemplate := template.Must(template.ParseFiles("auth/view/login.gohtml"))
		myTemplate.ExecuteTemplate(w, "login", map[string]interface{}{
			"err":   params.ByName("err"),
			"Title": "Cafe Mewa - Login",
		})
	}
}

func (Controller *LoginControllerImpl) LoginCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := model.UserLogin{
		Nama:     r.PostFormValue("nama"),
		Password: r.PostFormValue("password"),
	}
	serv, err := Controller.LoginService.Login(context.Background(), user)
	//fmt.Fprint(w, serv.Nama)
	if err != "" {
		myTemplate := template.Must(template.ParseFiles("auth/view/login.gohtml"))
		//w.WriteHeader(200)
		myTemplate.ExecuteTemplate(w, "login", map[string]interface{}{
			"error": err,
			"Title": "Cafe Mewa - Login",
		})
	} else {
		session, _ := config.Store.Get(r, config.SESSION_ID)

		session.Values["loggedIn"] = true
		session.Values["nama"] = serv.Nama
		session.Values["type"] = serv.Type

		session.Save(r, w)
		if serv.Type == "admin" {
			http.Redirect(w, r, "/pembelian/", http.StatusFound)
		} else if serv.Type == "kasir" {
			http.Redirect(w, r, "/pesanan/", http.StatusFound)
		}
	}
}

func (Controller *LoginControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/auth/login/", http.StatusFound)
}
