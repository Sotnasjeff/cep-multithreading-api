package dto

type GetViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type GetBrasilApiOutput struct {
	Cep     string `json:"cep"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Service string `json:"service"`
}
