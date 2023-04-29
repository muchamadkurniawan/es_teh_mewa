package main

import (
	_ "context"
	controllerAuth "eh_teh_mewa/auth/controller"
	repositoryAuth "eh_teh_mewa/auth/repository"
	serviceAuth "eh_teh_mewa/auth/service"
	"eh_teh_mewa/db"
	"eh_teh_mewa/master/controller"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/service"
	controllerPembelian "eh_teh_mewa/pembelian/controller"
	repositoryPembelian "eh_teh_mewa/pembelian/model/repository"
	servicePembelian "eh_teh_mewa/pembelian/service"
	controllerPesanan "eh_teh_mewa/pesanan/controller"
	repositoryPesanan "eh_teh_mewa/pesanan/model/repository"
	servicePesanan "eh_teh_mewa/pesanan/service"

	controllerProduk "eh_teh_mewa/produk/controller"
	repositoryProduk "eh_teh_mewa/produk/model/repository"
	serviceProduk "eh_teh_mewa/produk/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := db.GetConnect()
	UserRepoLogin := repositoryAuth.NewLoginRepository()
	UserServiceLogin := serviceAuth.NewLoginService(UserRepoLogin, db)
	UserControllerLogin := controllerAuth.NewLoginController(UserServiceLogin)

	userRepo := repository.NewUsersRepository()
	userService := service.NewUsersService(userRepo, db)
	usersController := controller.NewUsersController(userService)

	//
	pembelianRepository := repositoryPembelian.NewPembelianRepository()
	pembelianService := servicePembelian.NewPembelianService(pembelianRepository, db)
	pembelianController := controllerPembelian.NewPembelianController(pembelianService)

	//
	satuanRepository := repository.NewSatuanRepository()
	satuanService := service.NewSatuanService(satuanRepository, db)
	satuanController := controller.NewSatuanController(satuanService)

	//
	bahanBakuRepository := repository.NewBahanRepository()
	bahanBakuService := service.NewBahanbakuService(bahanBakuRepository, db)
	bahanBakuController := controller.NewBahanBakuController(bahanBakuService)

	//
	produkRepository := repositoryProduk.NewProdukRepo()
	produkService := serviceProduk.NewProdukService(produkRepository, db)
	produkController := controllerProduk.NewProdukController(produkService)

	//
	pesananRepository := repositoryPesanan.NewPesananRepository()
	pesananService := servicePesanan.NewPesananServie(pesananRepository, db)
	pesananController := controllerPesanan.NewPesananController(pesananService)

	//
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	router.ServeFiles("/storage/*filepath", http.Dir("./storage/"))

	router.GET("/auth/login/", UserControllerLogin.Login)
	router.POST("/auth/login/check/", UserControllerLogin.LoginCheck)
	router.POST("/auth/logout/", UserControllerLogin.Logout)

	router.GET("/user/", usersController.FindAll)
	router.GET("/user/create/", usersController.Create)
	router.GET("/user/register/", usersController.Register)
	router.POST("/user/store/", usersController.Store)
	router.GET("/user/show/:id/", usersController.FindById)
	router.POST("/user/update/:id/", usersController.Update)
	router.POST("/user/delete/:id/", usersController.Delete)

	router.GET("/satuan/", satuanController.FindAllSatuan)
	router.GET("/satuan/show/:id/", satuanController.FindById)
	router.POST("/satuan/store/", satuanController.Store)
	router.POST("/satuan/update/:id/", satuanController.Update)
	router.POST("/satuan/delete/:id/", satuanController.Delete)

	router.GET("/bahan-baku/", bahanBakuController.FindAllBahanBaku)
	router.GET("/bahan-baku/show/:id", bahanBakuController.FindByIdBahanBaku)
	router.POST("/bahan-baku/store/", bahanBakuController.Store)
	router.POST("/bahan-baku/update/:id/", bahanBakuController.Update)
	router.POST("/bahan-baku/delete/:id/", bahanBakuController.Delete)

	router.GET("/pembelian/", pembelianController.Index)
	router.GET("/pembelian/create/", pembelianController.Create)
	router.POST("/pembelian/store/", pembelianController.Store)
	router.GET("/pembelian/show/:id/", pembelianController.Show)
	router.POST("/pembelian/update/:id", pembelianController.Update)
	router.POST("/pembelian/delete/:id/", pembelianController.Delete)

	router.GET("/produk/", produkController.FindByAll)
	router.GET("/produk/show/:id/", produkController.FindById)
	router.GET("/produk/create/", produkController.Create)
	router.POST("/produk/store/", produkController.Store)
	router.POST("/produk/update/:id/", produkController.Update)
	router.POST("/produk/delete/:id/", produkController.Delete)

	router.GET("/pesanan/", pesananController.Index)
	router.GET("/pesanan/order/", pesananController.Create)
	router.POST("/pesanan/create/", pesananController.Store)
	router.GET("/pesanan/detail/:id/", pesananController.Show)
	router.GET("/pesanan/cetak/:id/", pesananController.Cetak)
	//router.POST("/pesanan/update/:id/", pesananController.Update)
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
