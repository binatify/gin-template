package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EnsureWithinMaxItems(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(EnsureWithinMaxItems(-1), DefaultMaxItems)
	assertion.Equal(EnsureWithinMaxItems(10000), DefaultMaxItems)
	assertion.Equal(EnsureWithinMaxItems(15), 15)
}
