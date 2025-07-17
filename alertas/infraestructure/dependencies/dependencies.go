package dependencies

import (
	"API-VITALVEST/alertas/application"
	"API-VITALVEST/alertas/infraestructure/http/controllers"
	"API-VITALVEST/alertas/infraestructure/http"
	"API-VITALVEST/core"
)

var (
	mySQL infraestructure.MYSQLRepository
)

func InitAlerta() {
	db := core.GetDBPool()
	if db == nil {
		panic("Error: la conexión a la base de datos es nil")
	}
	mySQL = *infraestructure.NewMYSQLRepository(db.DB)
}

func NewSaveAlertaController() *controllers.SaveAlertaController {
	useCase := application.NewSaveAlerta_UC(&mySQL)
	return controllers.NewSaveAlertaController(useCase)
}

func NewDeleteAlertaController() *controllers.DeleteAlertaController {
	useCase := application.NewDeleteAlerta(&mySQL)
	return controllers.NewDeleteAlertaController(useCase)
}

func NewUpdateAlertaController() *controllers.UpdateAlertaController {
	useCase := application.NewUpdateAlerta(&mySQL)
	return controllers.NewUpdateAlertaController(useCase)
}

func NewFindAllAlertaController() *controllers.GetAllAlertasController {
	useCase := application.NewGetAllAlertas_UC(&mySQL)
	return controllers.NewGetAllAlertasController(useCase)
}

func NewFindByIDAlertaController() *controllers.GetAlertaByIDController {
	useCase := application.NewGetAlertabyID(&mySQL)
	return controllers.NewGetAlertaByIDController(useCase)
}
