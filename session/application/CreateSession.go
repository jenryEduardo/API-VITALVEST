package application

import "API-VITALVEST/session/domain"

type Session struct {
	repo domain.Isesion
}

func NewSession(repo domain.Isesion)*Session{
	return &Session{repo: repo}
}


func (r *Session)Execute(data domain.Session)error{
	return r.repo.Save(data)
}
