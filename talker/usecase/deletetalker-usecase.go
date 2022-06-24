package talkerusecase

func (usecase *TalkersUseCase) DeleteTalker(id int) error {
	return usecase.Repo.DeleteTalker(id)
}
