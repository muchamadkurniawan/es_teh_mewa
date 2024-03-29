package main

import (
	_ "context"
	controllerAuth "es_teh_mewa/auth/controller"
	repositoryAuth "es_teh_mewa/auth/repository"
	serviceAuth "es_teh_mewa/auth/service"
	controllerBiaya "es_teh_mewa/biaya/controller"
	repositoryBiaya "es_teh_mewa/biaya/model/repository"
	serviceBiaya "es_teh_mewa/biaya/service"
	controllerDashboard "es_teh_mewa/dashboard/controller"
	repositoryDashboard "es_teh_mewa/dashboard/repository"
	serviceDashboard "es_teh_mewa/dashboard/service"
	"es_teh_mewa/db"
	"es_teh_mewa/master/controller"
	"es_teh_mewa/master/repository"
	"es_teh_mewa/master/service"
	controllerPembelian "es_teh_mewa/pembelian/controller"
	repositoryPembelian "es_teh_mewa/pembelian/model/repository"
	servicePembelian "es_teh_mewa/pembelian/service"
	controllerPesanan "es_teh_mewa/pesanan/controller"
	repositoryPesanan "es_teh_mewa/pesanan/model/repository"
	servicePesanan "es_teh_mewa/pesanan/service"
	controllerRekap "es_teh_mewa/rekap/controller"
	repositoryRekap "es_teh_mewa/rekap/model/repository"
	serviceRekap "es_teh_mewa/rekap/service"
	controllerStok "es_teh_mewa/stok/controller"
	repositoryStok "es_teh_mewa/stok/repository"
	serviceStok "es_teh_mewa/stok/service"

	controllerProduk "es_teh_mewa/produk/controller"
	repositoryProduk "es_teh_mewa/produk/model/repository"
	serviceProduk "es_teh_mewa/produk/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := db.GetConnect()

	StokRepository := repositoryStok.NewStokRepository()
	StokService := serviceStok.NewStokService(db, StokRepository)
	StokController := controllerStok.NewStokController(StokService)

	//
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

	rekapRepository := repositoryRekap.NewRekapRepository()
	rekapService := serviceRekap.NewRekapServiceImpl(rekapRepository, db)
	rekapController := controllerRekap.NewRekapController(rekapService)

	//
	biayaRepository := repositoryBiaya.NewBiayaRepository()
	biayaService := serviceBiaya.NewBiayaService(biayaRepository, db)
	biayaController := controllerBiaya.NewBiayaController(biayaService)
	//

	dashboardRepository := repositoryDashboard.NewDashboardRepository()
	dashboardService := serviceDashboard.NewDashboardService(dashboardRepository, db)
	dashboardController := controllerDashboard.NewDashboardController(dashboardService)

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

	router.GET("/rekap-kasir/", rekapController.Index)
	router.POST("/rekap-kasir/create/", rekapController.Create)

	router.GET("/biaya-kasir/", biayaController.Index)
	router.POST("/biaya-kasir/create/", biayaController.CreateBiaya)
	router.GET("/biaya-kasir/detail/:id/", biayaController.FindById)
	router.POST("/biaya-kasir/delete/:id/", biayaController.Delete)

	router.GET("/", dashboardController.GetRekap)
	router.GET("/cetak/:id/", dashboardController.Cetak)
	router.GET("/admin/detail-rekap/:id/", dashboardController.DetailRekap)
	router.GET("/admin/pesanan/detail/:id/", pesananController.ShowAdmin)
	router.GET("/admin/biaya/detail/:id/", biayaController.FindByIdAdmin)

	router.GET("/stok/", StokController.Index)
	router.GET("/detail-stok/", StokController.Detail)
	router.GET("/stock-kasir/", StokController.IndexKasir)
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
