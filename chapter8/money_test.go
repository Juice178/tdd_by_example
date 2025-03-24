package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert := assert.New(t)
	five, _ := GetMoney("Dollar", 5)
	assert.Equal(&Dollar{10}, five.Times(2))
	assert.Equal(&Dollar{15}, five.Times(3))
}

func TestEquality(t *testing.T) {
	assert := assert.New(t)
	five_dollar, _ := GetMoney("Dollar", 5)
	assert.True(five_dollar.Equals(Dollar{5}))
	assert.False(five_dollar.Equals(Dollar{6}))

	five_franc, _ := GetMoney("Franc", 5)

	assert.True(five_franc.Equals(Franc{5}))
	assert.False(five_franc.Equals(Franc{6}))
	assert.False(five_franc.Equals(five_dollar))
}

func TestFrancMultiplication(t *testing.T) {
	assert := assert.New(t)
	five := Franc{5}
	assert.Equal(&Franc{10}, five.Times(2))
	assert.Equal(&Franc{15}, five.Times(3))
}
