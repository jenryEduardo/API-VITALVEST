package domain

type Bme struct {
	Id int `json:"id"` 
	Temperatura_ambiente float64 `json:"temperatura_ambiente"`
	Humedad_relativa float64 `json:"humedad_relativa"`
}
