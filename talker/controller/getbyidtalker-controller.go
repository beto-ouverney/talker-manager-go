package talkercontroller

import (
	"encoding/json"
	"fmt"
)

// GetTalkerByID returns talker from Repository
func GetTalkerByID(controller *TalkersController, id int) (talkerJSON []byte, err error) {
	fmt.Println("CONTROLLER")
	talker, err := controller.ITalkersUseCase.GetTalkerByID(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("CONTROLLER2")
	fmt.Print(talker)
	talkerJSON, err = json.Marshal(talker)
	return
}
