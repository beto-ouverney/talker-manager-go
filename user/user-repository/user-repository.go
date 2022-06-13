package userrepository

import (
	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
	"github.com/google/uuid"
)

//UserRepository is the implementation of the user repository
type UserRepository struct {
	IUserRepository
}

//GetUserToken is a function that returns a token
func GetUserToken(user userentity.User) string {
	token := uuid.New().String()
	return token
}
