package handler

import "telefun/entity"

type CreatePhoneNumberResponse struct {
	PhoneNumber string `json:"phone_number"`
	Provider    string `json:"provider"`
}

func BuildCreatePhoneNumberResponse(pn entity.PhoneNumber) CreatePhoneNumberResponse {
	return CreatePhoneNumberResponse{
		PhoneNumber: pn.Get(),
		Provider:    pn.Provider(),
	}
}

type AvailablePhoneNumberResponse struct {
	Available bool `json:"available"`
}
