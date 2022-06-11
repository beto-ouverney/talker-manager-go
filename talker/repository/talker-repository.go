package talkerrepository

import (
	"encoding/json"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

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
