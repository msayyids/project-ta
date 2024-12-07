package repository

type LabaRugiRepository struct{}

type LabaRugiRepositoryInj interface{}

func NewLabaRugiRepository() LabaRugiRepositoryInj {
	return LabaRugiRepository{}
}
