package talkercontroller

//DeleteTalker is a function that deletes a talker from the list of talkers
func (controller *TalkersController) DeleteTalker(id int) (err error) {
	err = controller.UseCase.DeleteTalker(id)
	return
}
