package talkercontroller

import (
	"encoding/json"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

//AddTalker is a function that adds a new talker to the list of talkers
func (controller *TalkersController) AddTalker(newTalker *talker.Talker) (talkerJSON []byte, err error) {
	talker, err := controller.UseCase.AddTalker(newTalker)
	if err == nil {
		talkerJSON, err = json.Marshal(talker)
	}
	return

}
