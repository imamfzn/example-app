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

func (u *PhoneNumberUsecase) Create(ctx context.Context, number string) (entity.PhoneNumber, error) {
	pn, err := entity.ParsePhoneNumber(number)
	if err != nil {
		return entity.PhoneNumber{}, err
	}

	if err := u.repo.Create(ctx, pn); err != nil {
		return entity.PhoneNumber{}, err
	}
	return pn, nil
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
