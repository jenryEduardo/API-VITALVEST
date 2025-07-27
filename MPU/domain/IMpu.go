package domain

type IMpu interface {
	Save(Mpu Mpu) error
	FindAll() ([]Mpu, error)
	GetAll()([]Mpu,error)
	FindByID(id int) ([]Mpu,error)
	UpdateByID(id int, Mpu Mpu) error
	DeleteByID(id int) error
}