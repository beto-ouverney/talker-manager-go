package talkerusecase

import talker "github.com/beto-ouverney/talker-manager-go/talker/entity"

//EditTalkerUseCase is the implementation of the edit talker use case
func (useCase *TalkersUseCase) EditTalker(newTalker *talker.Talker) (*talker.Talker, error) {
	return useCase.Repo.EditTalker(newTalker)
}
