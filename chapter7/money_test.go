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

func TestFrancMultiplication(t *testing.T) {
	assert := assert.New(t)
	five := Franc{5}
	assert.Equal(Franc{10}, five.Times(2))
	assert.Equal(Franc{15}, five.Times(3))
}

func TestFrancEquality(t *testing.T) {
	assert := assert.New(t)
	var five Money = &Franc{5}
	assert.True(five.Equals(Franc{5}))
	assert.False(five.Equals(Franc{6}))
}
