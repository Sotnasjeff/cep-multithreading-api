package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	address, err := NewAddress("06364570")
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.NotEmpty(t, address.Cep)
	assert.Equal(t, "06364570", address.Cep)
}

func TestEnderecoWhenCepLengthIsLessThan8(t *testing.T) {
	address, err := NewAddress("063548")
	assert.Nil(t, address)
	assert.NotNil(t, err)
}

func TestEnderecoWhenCepIsEmptyString(t *testing.T) {
	address, err := NewAddress("")
	assert.Nil(t, address)
	assert.NotNil(t, err)
}
