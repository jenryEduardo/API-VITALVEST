package domain

type IGsr interface {
	Save(gsr Gsr) error
	FindAll() ([]Gsr, error)
	FindByID(id int) ([]Gsr,error)
	UpdateByID(id int, gsr Gsr) error
	DeleteByID(id int) error
}