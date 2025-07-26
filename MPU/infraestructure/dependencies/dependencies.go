package dependencies

import (
	"API-VITALVEST/MPU/application"
	"API-VITALVEST/MPU/infraestructure/http"
	"API-VITALVEST/MPU/infraestructure/http/controllers"
	"API-VITALVEST/core"
)

var (
	mySQL infraestructure.MYSQLRepository
)

func InitMPU() {
	db := core.GetDBPool()
	if db == nil {
		panic("Error: la conexión a la base de datos es nil")
	}
	mySQL = *infraestructure.NewMYSQLRepository(db.DB)
}

func NewSaveMPUController() *controllers.SaveMPUController {
	useCase := application.NewSaveMPU_UC(&mySQL)
	return controllers.NewSaveMPUController(useCase)
}

func NewDeleteMPUController() *controllers.DeleteMPUController {
	useCase := application.NewDeleteMPU(&mySQL)
	return controllers.NewDeleteMPUController(useCase)
}

func NewUpdateMPUController() *controllers.UpdateMPUController {
	useCase := application.NewUpdateMPU(&mySQL)
	return controllers.NewUpdateMPUController(useCase)
}

func NewFindAllMPUController() *controllers.GetAllMPUController {
	useCase := application.NewGetAllMPU_UC(&mySQL)
	return controllers.NewGetAllMPUController(useCase)
}

func NewGetAllTableMPUController() *controllers.GetAllTableMPU_Controller {
	useCase := application.NewGetMPU_UC(&mySQL)
	return controllers.NewGetAllTableMpuController(useCase)
}

func NewFindByIDMPUController() *controllers.GetMPUbyIDController {
	useCase := application.NewGetMPUbyID(&mySQL)
	return controllers.NewGetMPUbyIDController(useCase)
}
