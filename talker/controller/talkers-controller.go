package talkercontroller

import (
	talkerusecase "github.com/beto-ouverney/talker-manager-go/talker/usecase"
)

//TalkersController is the implementation of the talkers controller
type TalkersController struct {
	talkerusecase.ITalkersUseCase
}
