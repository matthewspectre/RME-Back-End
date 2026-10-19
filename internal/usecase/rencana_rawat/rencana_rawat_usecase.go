package rencanarawat

import (
	entity "rme/internal/entity/rencana_rawat"
	repo "rme/internal/repository/rencana_rawat"
)

type Usecase interface {
	Create(data *entity.RencanaRawat) error
	GetByID(id int) (*entity.RencanaRawat, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.RencanaRawat, error)
	Update(id int, updates map[string]interface{}) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.RencanaRawat) error {
	return u.repo.Create(data)
}

func (u *usecase) GetByID(id int) (*entity.RencanaRawat, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*entity.RencanaRawat, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}
