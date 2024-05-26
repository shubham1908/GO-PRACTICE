package pkgvet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintCheckVet(t *testing.T) {
	t.Run("testing with go vet", func(t *testing.T) {
		res := PrintCheckVet()
		assert.Equal(t, "ram", res)
	})

}
