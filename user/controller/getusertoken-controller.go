package usercontroller

import userentity "github.com/beto-ouverney/talker-manager-go/user/entity"

func (controller *UserController) GetUserToken(user userentity.User) string {
	token := controller.IUserCase.GetUserToken(user)
	return token
}
