package usercontroller

import userusecase "github.com/beto-ouverney/talker-manager-go/user/usecase"

//UserController is a controller for the UserUseCase
type UserController struct {
	userusecase.IUserUseCase
}
