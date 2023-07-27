package entity_test

import (
	"telefun/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePhoneNumber(t *testing.T) {
	type testcase struct {
		number           string
		expectedErr      error
		expectedNumber   string
		expectedProvider string
	}

	tests := map[string]testcase{
		"invalid region": {
			number:      "+61000000001",
			expectedErr: entity.ErrInvalidRegion,
		},
		"invalid format": {
			number:      "+62821OOOO1",
			expectedErr: entity.ErrInvalidFormat,
		},
		"invalid length less than 11": {
			number:      "+62821000001",
			expectedErr: entity.ErrInvalidLength,
		},
		"invalid length grater than 13": {
			number:      "+628210000000001",
			expectedErr: entity.ErrInvalidLength,
		},
		"invalid provider": {
			number:      "+628290000001",
			expectedErr: entity.ErrInvalidProvider,
		},
		"valid": {
			number:           "087823643331",
			expectedNumber:   "+6287823643331",
			expectedProvider: "xl",
		},
	}

	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			pn, err := entity.ParsePhoneNumber(tc.number)

			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedNumber, pn.Get())
			assert.Equal(t, tc.expectedNumber, pn.String())
			assert.Equal(t, tc.expectedProvider, pn.Provider())
			assert.Equal(t, tc.expectedErr == nil, pn.Valid())
		})
	}
}
