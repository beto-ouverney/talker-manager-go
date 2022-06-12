package talkerusecase

import (
	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

//GetAllTalkers returns all the talkers from Repository
func (useCase *TalkersUseCase) GetAllTalkers() (*[]talker.Talker, error) {
	talkers, err := useCase.Repo.GetAllTalkers()
	return talkers, err
}
