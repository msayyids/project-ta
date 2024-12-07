package service

import (
	"project-ta/repository"

	"github.com/jmoiron/sqlx"
)

type LabaRugiService struct {
	DB           sqlx.DB
	LabaRugiRepo repository.LabaRugiRepositoryInj
}

type LabaRugiServiceInj interface{}

func NewLabaRugiService(db sqlx.DB, lrp repository.LabaRugiRepositoryInj) LabaRugiServiceInj {
	return LabaRugiService{
		DB:           db,
		LabaRugiRepo: lrp,
	}
}
