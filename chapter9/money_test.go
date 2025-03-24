package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert := assert.New(t)
	five, _ := GetMoney("USD", 5)
	assert.Equal(&Dollar{10, "USD"}, five.Times(2))
	assert.Equal(&Dollar{15, "USD"}, five.Times(3))
}

func TestEquality(t *testing.T) {
	assert := assert.New(t)
	five_dollar, _ := GetMoney("USD", 5)
	assert.True(five_dollar.Equals(Dollar{5, "USD"}))
	assert.False(five_dollar.Equals(Dollar{6, "USD"}))

	five_franc, _ := GetMoney("CHF", 5)

	assert.True(five_franc.Equals(Franc{5, "CHF"}))
	assert.False(five_franc.Equals(Franc{6, "CHF"}))
	assert.False(five_franc.Equals(five_dollar))
}

func TestFrancMultiplication(t *testing.T) {
	assert := assert.New(t)
	five_franc, _ := GetMoney("CHF", 5)
	assert.Equal(&Franc{10, "CHF"}, five_franc.Times(2))
	assert.Equal(&Franc{15, "CHF"}, five_franc.Times(3))
}

func TestCurrency(t *testing.T) {
	assert := assert.New(t)
	five_dollar, _ := GetMoney("USD", 5)
	assert.Equal("USD", five_dollar.Currency())

	five_franc, _ := GetMoney("CHF", 5)
	assert.Equal("CHF", five_franc.Currency())

}
