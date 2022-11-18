package Multithreading

// {
// 	"cep": "96170-000",
// 	"logradouro": "",
// 	"complemento": "",
// 	"bairro": "",
// 	"localidade": "São Lourenço do Sul",
// 	"uf": "RS",
// 	"ibge": "4318804",
// 	"gia": "",
// 	"ddd": "53",
// 	"siafi": "8879"
// }

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// {"code":"96170-000","state":"RS","city":"São Lourenço do Sul","district":"","address":"","status":200,"ok":true,"statusText":"ok"}

type ConsultaAPIReponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}
