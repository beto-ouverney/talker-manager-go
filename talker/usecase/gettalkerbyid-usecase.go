package talkerusecase

import (
	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

//GetTalkerByID returns talker from Repository
func (useCase *TalkersUseCase) GetTalkerByID(id int) (*talker.Talker, error) {
	talker, err := useCase.Repo.GetTalkerByID(id)
	return talker, err
}
