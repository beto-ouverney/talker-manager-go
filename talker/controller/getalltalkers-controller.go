package talkercontroller

import (
	"encoding/json"
	"fmt"
)

//GetAllTalkers is a function that returns all the talkers from Repository
func (tC *TalkersController) GetAllTalkers() (talkerJSON []byte, err error) {
	talkers, err := tC.ITalkersUseCase.GetAllTalkers()
	if err != nil {
		panic(err)
	}
	for _, talker := range *talkers {
		fmt.Println(talker)
	}
	talkerJSON, err = json.Marshal(talkers)
	return
}
