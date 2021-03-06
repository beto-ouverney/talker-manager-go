package userintegration

import (
	usercontroller "github.com/beto-ouverney/talker-manager-go/user/controller"
	userrepository "github.com/beto-ouverney/talker-manager-go/user/repository"
	userusecase "github.com/beto-ouverney/talker-manager-go/user/usecase"
)

//UserIntegration is a function to use the user controller
func UserIntegration() *usercontroller.UserController {

	userRepository := &userrepository.UserRepository{}
	userUseCase := &userusecase.UserUseCase{userRepository}
	userController := &usercontroller.UserController{userUseCase}
	return userController
}
