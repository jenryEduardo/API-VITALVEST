package domain

type Ilogin interface{
	Login_app(login Login)([]Login,error)
}