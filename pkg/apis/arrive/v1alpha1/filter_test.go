package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	f := Filter{}
	assert.False(t, f.Match(nil))
}
