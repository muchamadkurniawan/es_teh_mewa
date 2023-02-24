package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/web"
)

type UsersServiceImpl struct {
	UsersRepo repository.UsersRepository
	DB        *sql.DB
}

func NewUsersService(usersRepository repository.UsersRepository, DB *sql.DB) UsersService {
	return &UsersServiceImpl{
		UsersRepo: usersRepository,
		DB:        DB,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)
	user := entity.Users{
		UserName:  request.Username,
		Password:  request.Password,
		Type_user: request.Type,
	}
	service.UsersRepo.InsertUsers(ctx, tx, user)
}

func (service *UsersServiceImpl) Update(ctx context.Context, response web.UsersResponse) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	user := entity.Users{
		Id:        int32(response.Id),
		UserName:  response.Username,
		Password:  response.Password,
		Type_user: response.Tipe,
	}
	service.UsersRepo.UpdateUsers(ctx, tx, user)
}

func (service *UsersServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)
	service.UsersRepo.DeleteUsers(context.Background(), tx, int32(id))
}

func (service *UsersServiceImpl) FindAll(ctx context.Context) []web.UsersResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)

	users, err := service.UsersRepo.FindByAllUsers(ctx, tx)
	if err != nil {
		panic(err)
	}
	usersResponse := helperMain.ToUserResponses(users)
	return usersResponse
}

func (service *UsersServiceImpl) FindById(ctx context.Context, id int) web.UsersResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helperMain.ErrorTx(tx)
	user, err := service.UsersRepo.FindByIdUsers(ctx, tx, int32(id))
	if err != nil {
		panic(err)
	}
	userResponse := helperMain.ToUserResponse(user)
	return userResponse
}
