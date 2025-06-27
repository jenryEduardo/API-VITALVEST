package application

import "API-VITALVEST/users/domain"

type GetAllUsers struct {
	repo domain.Iuser
}

func NewGetUsers(repo domain.Iuser)*GetAllUsers{
	return &GetAllUsers{repo: repo}
}

func(r *GetAllUsers)Execute()([]domain.User,error){
	return r.repo.Get()
}