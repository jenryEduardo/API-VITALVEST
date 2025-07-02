package domain

type Gsr struct {
	Id int `json:"id"` 
	Conductancia float64 `json:"conductancia"`
	Estado_hidratacion string `json:"estado_hidratacion"`
}
