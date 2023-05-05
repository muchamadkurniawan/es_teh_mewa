package service

import (
	"context"
	"database/sql"
	"es_teh_mewa/auth/model"
	"es_teh_mewa/auth/repository"
	"es_teh_mewa/helperMain"
	"golang.org/x/crypto/bcrypt"
)

type LoginServiceImpl struct {
	LoginRepo repository.LoginRepository
	DB        *sql.DB
}

func NewLoginService(loginRepository repository.LoginRepository, DB *sql.DB) LoginService {
	return &LoginServiceImpl{
		LoginRepo: loginRepository,
		DB:        DB,
	}
}

func (Service *LoginServiceImpl) Login(ctx context.Context, request model.UserLogin) (model.UserLogin, string) {
	tx, err := Service.DB.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	user := model.UserLogin{
		Nama:     request.Nama,
		Password: request.Password,
	}
	message := ""
	userRepo := Service.LoginRepo.Login(ctx, tx, user)
	if userRepo.Nama == "" {
		message = "Username Tidak ditemukan"
	} else {
		errPass := bcrypt.CompareHashAndPassword([]byte(userRepo.Password), []byte(request.Password))
		if errPass != nil {
			message = "Password Salah"
		}
	}
	return userRepo, message
}

func (Service *LoginServiceImpl) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
