package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContainer(t *testing.T) {
	c := NewContainer()

	assert.NotNil(t, c)
}

func TestContainer_ScheduledWorkload(t *testing.T) {
	c := NewContainer()

	assert.NotNil(t, c.ScheduledWorkload())
}

func TestContainer_ObservedWorkload(t *testing.T) {
	c := NewContainer()

	assert.NotNil(t, c.ObservedWorkload())
}
