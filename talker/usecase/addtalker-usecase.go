package talkerusecase

import talker "github.com/beto-ouverney/talker-manager-go/talker/entity"

//AddTalker is a implementation of the talkers use case
func (useCase *TalkersUseCase) AddTalker(newTalker *talker.Talker) (*talker.Talker, error) {
	return useCase.Repo.AddTalker(newTalker)
}
