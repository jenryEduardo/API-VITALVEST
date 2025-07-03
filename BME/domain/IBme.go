package domain

type IBme interface {
	Save(Bme Bme) error
	FindAll() ([]Bme, error)
	FindByID(id int) ([]Bme,error)
	UpdateByID(id int, Bme Bme) error
	DeleteByID(id int) error
}