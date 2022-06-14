package userrepository

import (
	"math/rand"
	"time"

	userentity "github.com/beto-ouverney/talker-manager-go/user/entity"
)

//IUserRepository interface for user repository
type IUserRepository interface {
	GetUserToken(user userentity.User) *string
}

//UserRepository is the implementation of the user repository
type UserRepository struct {
	IUserRepository
}

func tokenGenerator(numberChar int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, numberChar)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//GetUserToken is a function that returns a token
func (u *UserRepository) GetUserToken(user userentity.User) *string {
	var token *string
	tokenString := tokenGenerator(16)
	token = &tokenString
	return token
}
