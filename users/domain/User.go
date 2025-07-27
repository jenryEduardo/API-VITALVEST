package domain

type User struct {
	Id        int
	UserName  string `json:"username"`
	Passwords string `json:"passwords"`
}
