package repository

import (
	"telefun/entity"

	"github.com/jmoiron/sqlx"
)

type PhoneNumberRepository struct {
	db *sqlx.DB
}

func NewPhoneNumberRepository(db *sqlx.DB) *PhoneNumberRepository {
	return &PhoneNumberRepository{db}
}

func (r *PhoneNumberRepository) Create(pn entity.PhoneNumber) error {
	return nil
}

func (r *PhoneNumberRepository) Exists(pn entity.PhoneNumber) (bool, error) {
	return false, nil
}
