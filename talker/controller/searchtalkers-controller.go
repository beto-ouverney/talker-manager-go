package talkercontroller

import "encoding/json"

// SearchTalkers returns talkers from Repository
func (controller *TalkersController) SearchTalkers(searchTerm string) (talkersJSON []byte, err error) {
	talkers, err := controller.UseCase.SearchTalkers(searchTerm)
	if err == nil {
		talkersJSON, err = json.Marshal(talkers)
		return
	}
	return
}
