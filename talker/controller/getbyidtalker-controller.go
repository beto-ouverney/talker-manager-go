package talkercontroller

import (
	"encoding/json"
)

// GetTalkerByID returns talker from Repository
func (tC *TalkersController) GetTalkerByID(id int) (talkerJSON []byte, err error) {
	talker, err := tC.UseCase.GetTalkerByID(id)
	if err != nil {
		panic(err)
	}
	if talker != nil {
		talkerJSON, err = json.MarshalIndent(talker, "", "    ")
	}
	return
}
