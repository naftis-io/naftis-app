package source

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"testing"
)

func TestOs_Get(t *testing.T) {
	os := NewOs()

	list, err := os.Get()
	assert.NoError(t, err)
	assert.Len(t, list, 2)

	assert.Contains(t, list, label.OperatingSystem)
	assert.Contains(t, list, label.Architecture)
}
