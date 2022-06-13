package usercontroller

import userusecase "github.com/beto-ouverney/talker-manager-go/user/usecase"

type UserController struct {
	userusecase.IUserCase
}
