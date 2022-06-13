package userrepository

import userentity "github.com/beto-ouverney/talker-manager-go/user/entity"

//IUserRepository interface for user repository
type IUserRepository interface {
	GetUserToken(user userentity.User) string
}
