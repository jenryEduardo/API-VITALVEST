package domain

import "fmt"

const (
	MinPasswordLength = 8
	MaxPasswordLength = 16
)

type User struct {
	Id        int
	UserName  string `json:"username"`
	Passwords string `json:"passwords"`
	Id_chalecos int `json:"id_chalecos"`
}

func (l *User) ValidatePassword() error {
	length := len(l.Passwords)

	if length < MinPasswordLength {
		return fmt.Errorf("la contraseña debe tener al menos %d caracteres", MinPasswordLength)
	}

	if length > MaxPasswordLength {
		return fmt.Errorf("la contraseña no puede exceder %d caracteres", MaxPasswordLength)
	}

	return nil
}