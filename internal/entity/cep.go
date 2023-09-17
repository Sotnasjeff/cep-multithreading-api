package entity

import "errors"

var (
	ErrInvalidCep = errors.New("cep is invalid")
)

type Endereco struct {
	Cep string `json:"cep"`
}

func NewEndereco(cep string) (*Endereco, error) {
	endereco := &Endereco{
		Cep: cep,
	}

	err := endereco.Validate()
	if err != nil {
		return nil, err
	}

	return endereco, nil
}

func (e *Endereco) Validate() error {
	if e.Cep == "" || len(e.Cep) < 8 {
		return ErrInvalidCep
	}
	return nil
}
