package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEndereco(t *testing.T) {
	endereco, err := NewEndereco("06364570")
	assert.Nil(t, err)
	assert.NotNil(t, endereco)
	assert.NotEmpty(t, endereco.Cep)
	assert.Equal(t, "06364570", endereco.Cep)
}

func TestEnderecoWhenCepLengthIsLessThan8(t *testing.T) {
	endereco, err := NewEndereco("063548")
	assert.Nil(t, endereco)
	assert.NotNil(t, err)
}

func TestEnderecoWhenCepIsEmptyString(t *testing.T) {
	endereco, err := NewEndereco("")
	assert.Nil(t, endereco)
	assert.NotNil(t, err)
}
