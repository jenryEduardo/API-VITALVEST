package dependencies

import (
	"API-VITALVEST/MLX/application"
	"API-VITALVEST/MLX/infraestructure/http"
	"API-VITALVEST/MLX/infraestructure/http/controllers"
	"API-VITALVEST/core"
)

var (
	mySQL infraestructure.MYSQLRepository
)

func InitMLX() {
	db := core.GetDBPool()
	if db == nil {
		panic("Error: la conexión a la base de datos es nil")
	}
	mySQL = *infraestructure.NewMYSQLRepository(db.DB) 
}

func NewSaveMLXController() *controllers.SaveMLXController {
	useCase := application.NewSaveMLX_uc(&mySQL)
	return controllers.NewSaveMLXController(useCase)
}

func NewDeleteMLXController() *controllers.DeleteMLXController {
	useCase := application.NewDeleteMLX(&mySQL)
	return controllers.NewDeleteMLXController(useCase)
}

func NewUpdateMLXController() *controllers.UpdateMLXController {
	useCase := application.NewUpdateMLX(&mySQL)
	return controllers.NewUpdateMLXController(useCase)
}

func NewFindAllMLXController() *controllers.GetAllMLXController {
	useCase := application.NewGetAllMlx_UC(&mySQL)
	return controllers.NewGetAllMLXController(useCase)
}

func NewFindByIDMLXController() *controllers.GetMLXbyIDController {
	useCase := application.NewGetMlxbyID(&mySQL)
	return controllers.NewGetMLXbyIDController(useCase)
}
