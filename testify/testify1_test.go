package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TEst function
func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")
	// assert for nil (good for errors)
	assert.Nil(t, "nil")

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, "not nil") {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", "Something")

	}

}
