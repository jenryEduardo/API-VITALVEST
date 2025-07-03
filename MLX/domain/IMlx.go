package domain

type IMlx interface {
	Save(mlx Mlx) error
	FindAll() ([]Mlx, error)
	FindByID(id int) ([]Mlx,error)
	UpdateByID(id int, mlx Mlx) error
	DeleteByID(id int) error
}