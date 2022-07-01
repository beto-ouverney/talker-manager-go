package usercontroller

import (
	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
)

//GetUserToken is a function that returns a token for a user
func GetUserToken(controller *UserController, user userentity.User) *string {
	token := controller.IUserUseCase.GetUserToken(user)
	return token
}
