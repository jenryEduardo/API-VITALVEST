package dependencies

import (
	"API-VITALVEST/GSR/application"
	"API-VITALVEST/GSR/infraestructure/http"
	"API-VITALVEST/core/workerpool"
	"API-VITALVEST/GSR/infraestructure/http/controllers"
	"API-VITALVEST/core"
)

var (
	mySQL infraestructure.MYSQLRepository
	Pool *workerpool.WorkerPool
)

func InitGSR(pool *workerpool.WorkerPool) {
	Pool = pool
	db := core.GetDBPool()
	if db == nil {
		panic("Error: la conexión a la base de datos es nil")
	}
	mySQL = *infraestructure.NewMYSQLRepository(db.DB)
}

func NewSaveGSRController() *controllers.SaveGSRController {
	useCase := application.NewSaveGSR_UC(&mySQL)
	return controllers.NewSaveGSRController(useCase, Pool)
}

func NewDeleteGSRController() *controllers.DeleteGSRController {
	useCase := application.NewDeleteGSR(&mySQL)
	return controllers.NewDeleteGSRController(useCase)
}

func NewUpdateGSRController() *controllers.UpdateGSRController {
	useCase := application.NewUpdateGSR(&mySQL)
	return controllers.NewUpdateGSRController(useCase)
}

func NewFindAllGSRController() *controllers.GetAllGSRController {
	useCase := application.NewGetAllGsr_UC(&mySQL)
	return controllers.NewGetAllGSRController(useCase, Pool)
}

func NewFindByIDGSRController() *controllers.GetGSRbyIDController {
	useCase := application.NewGetGsrbyID(&mySQL)
	return controllers.NewGetGSRbyIDController(useCase)
}
