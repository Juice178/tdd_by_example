package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert := assert.New(t)
	five := Dollar{5}
	product := five.Times(2)
	assert.Equal(10, product.amount)
	product = five.Times(3)
	assert.Equal(15, product.amount)
}

func TestEquality(t *testing.T) {
	assert := assert.New(t)
	five := Dollar{5}
	assert.True(five.Equals(Dollar{5}))
	assert.False(five.Equals(Dollar{6}))
}
