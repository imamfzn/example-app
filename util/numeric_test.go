package util_test

import (
	"testing"

	"telefun/util"

	"github.com/stretchr/testify/assert"
)

func TestNumeric(t *testing.T) {
	assert.True(t, util.Numeric("123"))
	assert.False(t, util.Numeric("abc123"))
	assert.True(t, util.Numeric("0812"))
	assert.False(t, util.Numeric("0812o"))
	assert.False(t, util.Numeric("+628"))
}
