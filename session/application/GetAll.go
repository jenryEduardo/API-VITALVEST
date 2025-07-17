package application

import "API-VITALVEST/session/domain"

type GetAllData struct {
	repo domain.Isesion
}

func NewGetAll(repo domain.Isesion)*GetAllData{
	return &GetAllData{repo: repo}
}


func (r *GetAllData)Execute()([]domain.DataSession,error){
	return r.repo.GetAll()
}