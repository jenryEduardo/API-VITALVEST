package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/users/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type MYSQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMysqlRepo() *MYSQLRepository {
	conn := core.GetDBPool()
	return &MYSQLRepository{conn: conn}
}

func (r *MYSQLRepository) Save(user domain.User) error {

	query := "INSERT INTO users(username,passwords) VALUES(?,?)"

	hash, errores := bcrypt.GenerateFromPassword([]byte(user.Passwords), bcrypt.DefaultCost)

	if errores != nil {
		fmt.Print("no se pudo realizar el hash")
	}

	_, err := r.conn.DB.Exec(query, &user.UserName, &hash)
	if err != nil {
		return err
	}

	return nil

}

func (r *MYSQLRepository) Delete(id int) error {

	query := "DELETE FROM users WHERE id = ?"
	_, err := r.conn.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar al usuario: %w", err)
	}

	return nil
}

func (r *MYSQLRepository) Update(user domain.User, id int) error {
	query := "UPDATE users SET username=?,passwords=? WHERE id = ?"

	response, err := r.conn.DB.Exec(query, &user.UserName, &user.Passwords, id)

	if err != nil {
		fmt.Println("no se pudo actualizar el dato verifique el sinstaxis o los datos")
	}

	rows, _ := response.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("error no se actualizo ningun dato")
	}

	return err
}

func (r *MYSQLRepository) Get() ([]domain.User, error) {
	query := "SELECT id,username FROM users"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.UserName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Verificamos si hubo errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MYSQLRepository) Login(name string) (*domain.User, error) {
	query := "SELECT id, username FROM users WHERE username = ? LIMIT 1"
	row := r.conn.DB.QueryRow(query, name)

	var user domain.User
	err := row.Scan(&user.Id, &user.UserName)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado: %w", err)
	}

	return &user, nil
}
