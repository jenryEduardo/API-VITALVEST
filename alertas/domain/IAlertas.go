package domain

type IAlertas interface {
	Save(Alertas Alerta) error
	FindAll() ([]Alerta, error)
	FindByID(id int) ([]Alerta, error)
	UpdateByID(id int, Alertas Alerta) error
	DeleteByID(id int) error
}