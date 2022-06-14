package userusecase

import (
	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
)

//GetUserToken is a function that returns a token for a user
func (userUC *UserUseCase) GetUserToken(user userentity.User) *string {

	var token *string
	token = userUC.URepo.GetUserToken(user)

	return token
}
