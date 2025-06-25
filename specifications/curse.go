package specifications

import (
	"testing"

	"github.com/alecthomas/assert"
)

type Curser interface {
	Curse(string) (string, error)
}

func CurseSpecification(t testing.TB, curser Curser) {
	got, err := curser.Curse("John")
	assert.NoError(t, err)
	assert.Equal(t, got, "Crazy John")
}
