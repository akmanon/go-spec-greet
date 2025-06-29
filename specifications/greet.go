package specifications

import (
	"testing"

	"github.com/alecthomas/assert"
)

type Greeter interface {
	Greet(string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Mike")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello Mike")
}
