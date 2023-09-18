package entity

import "errors"

var (
	ErrInvalidCep = errors.New("cep is invalid")
)

type Address struct {
	Cep string `json:"cep"`
}

func NewAddress(cep string) (*Address, error) {
	address := &Address{
		Cep: cep,
	}

	err := address.Validate()
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (e *Address) Validate() error {
	if e.Cep == "" || len(e.Cep) < 8 {
		return ErrInvalidCep
	}
	return nil
}
