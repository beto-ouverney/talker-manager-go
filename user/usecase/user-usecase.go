package userusecase

import (
	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
	userrepository "github.com/beto-ouverney/talker-manager-go/user/user-repository"
)

//IUserUseCase interface for user usecase
type IUserCase interface {
	GetUserToken(user userentity.User) string
}

//UserCase struct for user usecase
type UserUseCase struct {
	Repo userrepository.IUserRepository
}
