package domain


type Isesion interface{
	Save(sesion Session)error
	GetAll()([]Session,error)
}
