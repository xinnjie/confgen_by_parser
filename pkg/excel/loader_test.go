package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoader_Load(t *testing.T) {
	loader, err := NewLoader("../../testdata/basic-field.xlsx")
	assert.NoError(t, err)
	assert.NoError(t, loader.Load("Basic"))
}
