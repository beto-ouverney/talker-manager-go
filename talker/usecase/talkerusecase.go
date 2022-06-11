package talkerusecase

import (
	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
	talkerrepository "github.com/beto-ouverney/talker-manager-go/talker/repository"
)

//ITalkersUseCase is the interface for the talkers use case
type ITalkersUseCase interface {
	GetAllTalkers() (*[]talker.Talker, error)
}

//TalkersUseCase is the implementation of the talkers use case
type TalkersUseCase struct {
	Irepo talkerrepository.ITalkerRepository
}
