package talkerusecase

import talker "github.com/beto-ouverney/talker-manager-go/talker/entity"

//SearchTalkers is a implementation of the talkers use case
func (useCase *TalkersUseCase) SearchTalkers(search string) (*[]talker.Talker, error) {
	return useCase.Repo.SearchTalkers(search)
}
