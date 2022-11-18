package Multithreading

type Consulta struct {
	Buscador   string `json:"buscador"`
	CEP        string `json:"cep"`
	UF         string `json:"uf"`
	Cidade     string `json:"cidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}
