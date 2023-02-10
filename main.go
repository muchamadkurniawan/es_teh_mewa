package main

import (
	_ "context"
	"eh_teh_mewa/db"
	"eh_teh_mewa/master/controller"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := db.GetConnect()
	userRepo := repository.NewUsersRepository()
	userService := service.NewUsersService(userRepo, db)
	usersController := controller.NewUsersController(userService)

	//
	//pembelianRepository := repositoryPembelian.NewPembelianRepository()
	//pembelianService := servicePembelian.NewPembelianService(pembelianRepository, db)
	//pembelianController := controllerPembelian.NewPembelianController(pembelianService)

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	router.GET("/user/", usersController.FindAll)
	router.GET("/user/create/", usersController.Create)

	//router.GET("/pembelian/", pembelianController.Index)
	//router.GET("/pembelian/create/", pembelianController.Create)
	//router.POST("/pembelian/store/", pembelianController.Store)
	//router.GET("/pembelian/show/:id/", pembelianController.Show)
	//router.POST("/pembelian/delete/:id/", pembelianController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	fmt.Println("Running Program")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
