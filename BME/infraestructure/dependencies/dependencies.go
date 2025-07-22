package dependencies

import (
	"API-VITALVEST/BME/application"
	"API-VITALVEST/BME/infraestructure/http"
	"API-VITALVEST/BME/infraestructure/http/controllers"
	"API-VITALVEST/core"
)

var (
	mySQL infraestructure.MYSQLRepository
)

func InitBME() {
	db := core.GetDBPool()
	if db == nil {
		panic("Error: la conexión a la base de datos es nil")
	}
	mySQL = *infraestructure.NewMYSQLRepository(db.DB)
}

func NewSaveBMEController() *controllers.SaveBMEController {
	useCase := application.NewSaveBME_UC(&mySQL)
	return controllers.NewSaveBMEController(useCase)
}

func NewDeleteBMEController() *controllers.DeleteBMEController {
	useCase := application.NewDeleteBME(&mySQL)
	return controllers.NewDeleteBMEController(useCase)
}

func NewUpdateBMEController() *controllers.UpdateBMEController {
	useCase := application.NewUpdateBME(&mySQL)
	return controllers.NewUpdateBMEController(useCase)
}

func NewFindAllBMEController() *controllers.GetAllBMEController {
	useCase := application.NewGetAllBME_UC(&mySQL)
	return controllers.NewGetAllBMEController(useCase)
}

func NewFindByIDBMEController() *controllers.GetBMEbyIDController {
	useCase := application.NewGetBMEbyID(&mySQL)
	return controllers.NewGetBMEbyIDController(useCase)
}
