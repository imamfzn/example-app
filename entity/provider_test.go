package entity_test

import (
	"telefun/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	type testcase struct {
		number   string
		expected string
	}

	tests := []testcase{
		{number: "+628520000001", expected: "telkomsel"},
		{number: "+628110000001", expected: "telkomsel"},
		{number: "+628210000001", expected: "telkomsel"},
		{number: "+628780000001", expected: "xl"},
		{number: "+628170000001", expected: "xl"},
		{number: "+628960000001", expected: "tri"},
	}

	for _, tc := range tests {
		t.Run(tc.expected, func(t *testing.T) {
			assert.Equal(t, tc.expected, entity.Provider(tc.number))
		})
	}
}
