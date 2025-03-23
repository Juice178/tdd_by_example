package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	five.Times(2)
	assert := assert.New(t)
	assert.Equal(10, five.amount)
}
