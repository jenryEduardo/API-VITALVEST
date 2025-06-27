package domain


type Isesion interface{
	Save(sesion Sesion)error
	Delete(id int)error
	Update(sesion Sesion,id int)error
	GetAll()([]Sesion,error)
	GetById(id int)([]Sesion,error)
}
