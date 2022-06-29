package talkercontroller

import (
	"encoding/json"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

// EditTalker is a function that edits a talker
func (controller *TalkersController) EditTalker(newTalker *talker.Talker) (talkerJSON []byte, err error) {
	talker, err := controller.UseCase.EditTalker(newTalker)
	if err == nil {
		talkerJSON, err = json.Marshal(talker)
	}
	return
}
