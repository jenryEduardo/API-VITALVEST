package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/login/domain"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Mysql_conn struct {
	conn *core.Conn_MySQL
}

func NewMysqlRepo() *Mysql_conn {
	conn := core.GetDBPool()
	return &Mysql_conn{conn: conn}
}

func (r *Mysql_conn) Login_app(login domain.Login) ([]domain.Login, error) {
	query := "SELECT username, passwords FROM users WHERE username = ?"

	rows, err := r.conn.DB.Query(query, login.UserName)
	if err != nil {
		fmt.Println("Error en query:", err)
		return nil, fmt.Errorf("error al consultar la base de datos")
	}
	defer rows.Close()

	var data []domain.Login

	for rows.Next() {
		var dbUser domain.Login
		err := rows.Scan(&dbUser.UserName, &dbUser.Password)
		if err != nil {
			fmt.Println("Error al escanear:", err)
			continue
		}

		// Verificar la contraseña
		if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(login.Password)) != nil {
			fmt.Println("Contraseña incorrecta para usuario:", login.UserName)
			return nil, fmt.Errorf("credenciales incorrectas")
		}

		// No incluir la contraseña en la respuesta
		dbUser.Password = ""
		data = append(data, dbUser)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	return data, nil
}
