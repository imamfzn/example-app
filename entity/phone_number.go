package entity

import (
	"errors"
	"strings"
	"telefun/util"
)

const (
	MinNumberLength = 11
	MaxNumberLength = 13
)

var (
	ErrInvalidProvider = errors.New("Provider is not recognized")
	ErrInvalidLength   = errors.New("Invalid phone number length")
	ErrInvalidRegion   = errors.New("Region is not supported")
	ErrInvalidFormat   = errors.New("Invalid phone number format")
)

var cutPrefixes = []string{"08", "628", "+628"}

type PhoneNumber struct {
	number   string
	provider string
	valid    bool
}

// 08...
// 628...
// +628....
func ParsePhoneNumber(number string) (PhoneNumber, error) {
	var found bool
	for _, prefix := range cutPrefixes {
		number, found = strings.CutPrefix(number, prefix)
		if found {
			break
		}
	}

	if !found {
		return PhoneNumber{}, ErrInvalidRegion
	}

	if !util.Numeric(number) {
		return PhoneNumber{}, ErrInvalidFormat
	}

	// trimed number + "08"
	length := len(number) + 2
	if length < MinNumberLength || length > MaxNumberLength {
		return PhoneNumber{}, ErrInvalidLength
	}

	// convert into E164 format
	number = "+628" + number
	provider := Provider(number)
	if provider == "" {
		return PhoneNumber{}, ErrInvalidProvider
	}

	return PhoneNumber{number: number, provider: provider, valid: true}, nil

}

func (p PhoneNumber) String() string {
	return p.number
}

func (p PhoneNumber) Get() string {
	return p.number
}

func (p PhoneNumber) Valid() bool {
	return p.valid
}

func (p PhoneNumber) Provider() string {
	return p.provider
}
