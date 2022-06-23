package talkercontroller

import (
	"encoding/json"
)

// GetTalkerByID returns talker from Repository
func GetTalkerByID(controller *TalkersController, id int) (talkerJSON []byte, err error) {
	talker, err := controller.ITalkersUseCase.GetTalkerByID(id)
	if err != nil {
		panic(err)
	}
	talkerJSON, err = json.Marshal(talker)
	return
}
