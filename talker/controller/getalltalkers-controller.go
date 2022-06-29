package talkercontroller

import (
	"encoding/json"
)

//GetAllTalkers is a function that returns all the talkers from Repository
func (tC *TalkersController) GetAllTalkers() (talkerJSON []byte, err error) {
	talkers, err := tC.UseCase.GetAllTalkers()
	if err != nil {
		panic(err)
	}
	talkerJSON, err = json.MarshalIndent(talkers, "", "    ")
	return
}
