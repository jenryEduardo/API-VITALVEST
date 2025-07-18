package infraestructure

import (
	"API-VITALVEST/MPU/domain"
	"database/sql"
	"fmt"
	"log"
)

type MYSQLRepository struct {
	db *sql.DB
}

func NewMYSQLRepository(db *sql.DB) *MYSQLRepository {
	return &MYSQLRepository{db: db}
}

// Guarda un nuevo registro en la tabla mpu
func (r *MYSQLRepository) Save(mpu domain.Mpu) error {
	query := `INSERT INTO mpu (aceleracion_x, aceleracion_y, aceleracion_z, pasos, nivel_actividad) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query,
		mpu.Mpu6050.Aceleracion.X,
		mpu.Mpu6050.Aceleracion.Y,
		mpu.Mpu6050.Aceleracion.Z,
		mpu.Pasos,
		mpu.NivelActividad,
	)
	return err
}

// Actualiza un registro por id
func (r *MYSQLRepository) UpdateByID(id int, mpu domain.Mpu) error {
	query := `UPDATE mpu SET aceleracion_x=?, aceleracion_y=?, aceleracion_z=?, pasos=?, nivel_actividad=? WHERE id = ?`
	result, err := r.db.Exec(query,
		mpu.Mpu6050.Aceleracion.X,
		mpu.Mpu6050.Aceleracion.Y,
		mpu.Mpu6050.Aceleracion.Z,
		mpu.Pasos,
		mpu.NivelActividad,
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
	query := `SELECT id, aceleracion_x, aceleracion_y, aceleracion_z, pasos, nivel_actividad FROM mpu`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var m domain.Mpu
		err := rows.Scan(&m.Id, &m.Mpu6050.Aceleracion.X, &m.Mpu6050.Aceleracion.Y, &m.Mpu6050.Aceleracion.Z, &m.Pasos, &m.NivelActividad)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, m)
	}

	return MPUs, rows.Err()
}

// Obtiene registros por id (aunque suele ser solo uno, retorna slice por consistencia)
func (r *MYSQLRepository) FindByID(id int) ([]domain.Mpu, error) {
	query := `SELECT id, aceleracion_x, aceleracion_y, aceleracion_z, pasos, nivel_actividad FROM mpu WHERE id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var m domain.Mpu
		err := rows.Scan(&m.Id, &m.Mpu6050.Aceleracion.X, &m.Mpu6050.Aceleracion.Y, &m.Mpu6050.Aceleracion.Z, &m.Pasos, &m.NivelActividad)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, m)
	}

	return MPUs, rows.Err()
}
