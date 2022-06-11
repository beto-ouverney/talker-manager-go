package talkerintegration

import (
	talkercontroller "github.com/beto-ouverney/talker-manager-go/talker/controller"
	talkerrepository "github.com/beto-ouverney/talker-manager-go/talker/repository"
	talkerusecase "github.com/beto-ouverney/talker-manager-go/talker/usecase"
)

//TalkersIntegration is the implementation of the talkers integration
func TalkersIntegration() *talkercontroller.TalkersController {
	talkerRepository := &talkerrepository.TalkerRepository{}
	talkersUseCase := &talkerusecase.TalkersUseCase{talkerRepository}
	talkersController := &talkercontroller.TalkersController{talkersUseCase}
	return talkersController
}
