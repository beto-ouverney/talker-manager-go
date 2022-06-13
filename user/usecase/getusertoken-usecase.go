package userusecase

import userentity "github.com/beto-ouverney/talker-manager-go/user/entity"

func (usesUC *UserUseCase) GetUserToken(user userentity.User) string {
	token := usesUC.Repo.GetUserToken(user)
	return token
}
