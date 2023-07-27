package repository

import (
	"telefun/entity"

	"github.com/jmoiron/sqlx"
)

type phoneNumberRepository struct {
	db *sqlx.DB
}

func NewPhoneNumberRepository(db *sqlx.DB) *phoneNumberRepository {
	return &phoneNumberRepository{db}
}

func (r *phoneNumberRepository) Create(pn entity.PhoneNumber) error {
	return nil
}

func (r *phoneNumberRepository) Exists(pn entity.PhoneNumber) (bool, error) {
	return false, nil
}
