package domain


type Isesion interface{
	Save(sesion Session)error
	GetAll()([]DataSession,error)
}
