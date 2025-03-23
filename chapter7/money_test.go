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

	assert.True((&Dollar{5}).Equals(Dollar{5}))
	assert.False((&Dollar{5}).Equals(Dollar{6}))

	assert.True((&Franc{5}).Equals(Franc{5}))
	assert.False((&Franc{5}).Equals(Franc{6}))

	assert.True((&Dollar{5}).Equals(Dollar{5}))

	assert.False((&Franc{5}).Equals(Dollar{5}))
}

func TestFrancMultiplication(t *testing.T) {
	assert := assert.New(t)
	five := Franc{5}
	assert.Equal(Franc{10}, five.Times(2))
	assert.Equal(Franc{15}, five.Times(3))
}
