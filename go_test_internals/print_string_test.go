package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	// t.Run("test print string, assertion failed", func(t *testing.T) {
	// 	res, err := funcToTest(false)
	// 	assert.Equal(t, res, "")
	// 	assert.Equal(t, err, nil)
	// })

	t.Run("test print string, assertion success", func(t *testing.T) {
		res, err := funcToTest(false)
		assert.Equal(t, res, "test passed")
		assert.Equal(t, err, nil)
	})
}
