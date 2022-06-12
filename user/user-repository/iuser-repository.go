package userrepository

//IUserRepository interface for user repository
type IUserRepository interface {
	GetToken(user entity.User) string
}
