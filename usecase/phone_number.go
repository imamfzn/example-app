package usecase

import (
	"context"
	"telefun/entity"
	"telefun/repository"
)

type PhoneNumberUsecase struct {
	repo *repository.PhoneNumberRepository
}

func NewPhoneNumberUsecase(repo *repository.PhoneNumberRepository) *PhoneNumberUsecase {
	return &PhoneNumberUsecase{repo}
}

func (u *PhoneNumberUsecase) Create(number string) (entity.PhoneNumber, error) {
	return entity.PhoneNumber{}, nil
}

func (u *PhoneNumberUsecase) Available(ctx context.Context, number string) (bool, error) {
	pn, err := entity.ParsePhoneNumber(number)
	if err != nil {
		return false, err
	}

	exists, err := u.repo.Exists(ctx, pn)
	if err != nil {
		return false, err
	}
	return !exists, nil
}
