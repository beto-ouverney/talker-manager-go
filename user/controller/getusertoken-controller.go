package usercontroller

import (
	"fmt"

	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
)

//GetUserToken is a function that returns a token for a user
func GetUserToken(controller *UserController, user userentity.User) *string {
	fmt.Println("CONTROLLER")
	token := controller.IUserUseCase.GetUserToken(user)
	fmt.Println("CONTROLLER2")
	fmt.Println(token)
	return token
}
