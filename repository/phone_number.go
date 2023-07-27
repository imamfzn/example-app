package repository

import (
	"context"
	"database/sql"
	"errors"
	"telefun/entity"
	"telefun/util"

	"github.com/jmoiron/sqlx"
)

type PhoneNumberRepository struct {
	db *sqlx.DB
}

func NewPhoneNumberRepository(db *sqlx.DB) *PhoneNumberRepository {
	return &PhoneNumberRepository{db}
}

func (r *PhoneNumberRepository) Create(ctx context.Context, pn entity.PhoneNumber) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO phone_numbers (number, provider) VALUES ($1, $2)", pn.Get(), pn.Provider())
	if err != nil {
		if util.PostgresUniqueViolation(err) {
			return entity.ErrPhoneNumberTaken
		}
		return err
	}
	return nil
}

func (r *PhoneNumberRepository) Exists(ctx context.Context, pn entity.PhoneNumber) (bool, error) {
	var res int

	err := r.db.GetContext(ctx, &res, "SELECT 1 FROM phone_numbers WHERE number = $1 LIMIT 1", pn.Get())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
