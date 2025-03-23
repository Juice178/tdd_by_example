package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert := assert.New(t)
	five := Dollar{5}
	assert.Equal(Dollar{10}, five.Times(2))
	assert.Equal(Dollar{15}, five.Times(3))
}

func TestEquality(t *testing.T) {
	assert := assert.New(t)
	// var five Money
	var five Money = &Dollar{5}
	assert.True(five.Equals(Dollar{5}))
	assert.False(five.Equals(Dollar{6}))
}
