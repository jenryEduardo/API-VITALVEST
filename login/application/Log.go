package application

import "API-VITALVEST/login/domain"

type Loginapp struct {
	repo domain.Ilogin
}

func NewLogin(repo domain.Ilogin)*Loginapp{
	return &Loginapp{repo: repo}
}

func(c *Loginapp)Execute(login domain.Login)([]domain.Login,error){
	return c.repo.Login_app(login)
}
