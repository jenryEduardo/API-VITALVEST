package domain

type Bme struct {
	Id int `json:"id"` 
	Temperatura float64 `json:"temperatura"`
	Presion float64 `json:"presion"`
	Humedad float64 `json:"humedad"`
}
