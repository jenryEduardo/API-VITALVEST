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
	query := "INSERT INTO users(username, passwords) VALUES (?, ?)"

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Passwords), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("no se pudo hashear la contraseña: %w", err)
	}

	_, err = r.conn.DB.Exec(query, user.UserName, hash)
	if err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
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
	query := "UPDATE users SET username = ?, passwords = ? WHERE id = ?"

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Passwords), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("no se pudo hashear la nueva contraseña: %w", err)
	}

	result, err := r.conn.DB.Exec(query, user.UserName, hash, id)
	if err != nil {
		return fmt.Errorf("no se pudo actualizar el usuario: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no se actualizó ningún usuario")
	}

	return nil
}


func (r *MYSQLRepository) Get() ([]domain.User, error) {
	query := "SELECT id, username FROM users"
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func (r *MYSQLRepository) Login(username, password string) (*domain.User, error) {
	query := "SELECT id, username, passwords FROM users WHERE username = ? LIMIT 1"
	row := r.conn.DB.QueryRow(query, username)

	var user domain.User
	err := row.Scan(&user.Id, &user.UserName, &user.Passwords)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Passwords), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("contraseña incorrecta")
	}

	return &user, nil
}

