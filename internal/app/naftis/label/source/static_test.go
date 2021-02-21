package source

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"testing"
)

func TestStatic_Get(t *testing.T) {
	static := NewStatic(label.SourceLabels{
		"foo.bar": "example",
	})

	list, err := static.Get()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	assert.Contains(t, list, "foo.bar")
}
