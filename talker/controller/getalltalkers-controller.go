package talkercontroller

import (
	"encoding/json"
)

//GetAllTalkers is a function that returns all the talkers from Repository
func (tC *TalkersController) GetAllTalkers() (talkerJSON []byte, err error) {
	talkers, err := tC.ITalkersUseCase.GetAllTalkers()
	if err != nil {
		panic(err)
	}
	talkerJSON, err = json.Marshal(talkers)
	return
}
