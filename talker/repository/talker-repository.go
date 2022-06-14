package talkerrepository

import (
	"encoding/json"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

//ITalkerRepository is the interface for the talker repository
type ITalkerRepository interface {
	GetAllTalkers() (*[]talker.Talker, error)
	GetTalkerByID(id int) (*talker.Talker, error)
}

//TalkerRepository is the implementation of the talker repository
type TalkerRepository struct {
	ITalkerRepository
}

//GetAllTalkers is a function that returns all the talkers
func (t *TalkerRepository) GetAllTalkers() (*[]talker.Talker, error) {
	jsonFile, err := readJSON()
	if err != nil {
		return nil, err
	}
	var talkers *[]talker.Talker
	err = json.Unmarshal(jsonFile, &talkers)
	return talkers, err
}

//GetTalkerByID is a function that returns a talker by id
func (t *TalkerRepository) GetTalkerByID(id int) (*talker.Talker, error) {
	jsonFile, err := readJSON()
	if err != nil {
		return nil, err
	}
	var talkers []talker.Talker
	err = json.Unmarshal(jsonFile, &talkers)
	var talker *talker.Talker
	for _, v := range talkers {
		if v.ID == id {
			talker = &v
		}
	}
	return talker, err
}
