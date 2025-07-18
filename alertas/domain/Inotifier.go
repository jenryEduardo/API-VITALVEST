package domain

type Notifier interface {
	EnviarAlerta(alerta Alerta) error
}
