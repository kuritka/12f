package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testee struct {
	Id int
	Value string
}

func TestParseNil(t *testing.T) {
	err := Parse(nil)
	assert.Error(t, err)
}

func TestParseNotPointer(t *testing.T) {
	testee := Testee{}
	err := Parse(testee)
	assert.Error(t, err)
}
