package userusecase

import (
	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
	userrepository "github.com/beto-ouverney/talker-manager-go/user/repository"
)

//IUserUseCase interface for user usecase
type IUserUseCase interface {
	GetUserToken(user userentity.User) *string
}

//UserUseCase struct for user usecase
type UserUseCase struct {
	URepo userrepository.IUserRepository
}
