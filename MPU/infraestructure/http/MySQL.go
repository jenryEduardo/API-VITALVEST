package infraestructure

import (
	"API-VITALVEST/MPU/domain"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type MYSQLRepository struct {
	db *sql.DB
}

func NewMYSQLRepository(db *sql.DB) *MYSQLRepository {
	return &MYSQLRepository{db: db}
}

// Guarda un nuevo registro en la tabla mpu
func (r *MYSQLRepository) Save(mpu domain.Mpu) error {
	tiempo := time.Now()
	query := `INSERT INTO mpu (pasos,fecha) VALUES (?, ?)`
	_, err := r.db.Exec(query,
		&mpu.Pasos,
		tiempo,
	)
	return err
}

// Actualiza un registro por id
func (r *MYSQLRepository) UpdateByID(id int, mpu domain.Mpu) error {
	query := `UPDATE mpu SET pasos=? WHERE id = ?`
	result, err := r.db.Exec(query,
		mpu.Pasos,
		id,
	)

	if err != nil {
		log.Println("No se pudo actualizar el dato, verifique la sintaxis o los datos:", err)
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("error: no se actualizó ningún dato")
	}

	return nil
}

// Elimina un registro por id
func (r *MYSQLRepository) DeleteByID(id int) error {
	query := `DELETE FROM mpu WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("No se pudo eliminar el sensor MPU, verifique el id o la sintaxis SQL:", err)
		return err
	}
	return nil
}

// Obtiene todos los registros
func (r *MYSQLRepository) FindAll() ([]domain.Mpu, error) {
	query := `SELECT id, pasos,fecha FROM mpu`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var m domain.Mpu
		err := rows.Scan(&m.Id, &m.Pasos,&m.Fecha)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, m)
	}

	return MPUs, rows.Err()
}

// Obtiene registros por id (aunque suele ser solo uno, retorna slice por consistencia)
func (r *MYSQLRepository) FindByID(id int) ([]domain.Mpu, error) {
	query := `SELECT id, pasos, fecha FROM mpu WHERE id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var m domain.Mpu
		err := rows.Scan(&m.Id, &m.Pasos,&m.Fecha)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, m)
	}

	return MPUs, rows.Err()
}
