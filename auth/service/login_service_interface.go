package service

import (
	"context"
	"es_teh_mewa/auth/model"
)

type LoginService interface {
	Login(ctx context.Context, request model.UserLogin) (model.UserLogin, string)
	Logout(ctx context.Context) error
}
