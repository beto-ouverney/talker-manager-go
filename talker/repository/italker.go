package talkerrepository

import talker "github.com/beto-ouverney/talker-manager-go/talker/entity"

//ITalkerRepository is the interface for the talker repository
type ITalkerRepository interface {
	GetAllTalkers() (*[]talker.Talker, error)
	GetTalkerByID(id int) (*talker.Talker, error)
}
