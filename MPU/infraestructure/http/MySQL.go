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

func (r *MYSQLRepository) Save(MPU domain.Mpu) error {
	query := "INSERT INTO mpu (aceleracion_x, aceleracion_y, aceleracion_z, pasos, nivel_actividad) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, MPU.Aceleracion_x, MPU.Aceleracion_y, MPU.Aceleracion_z, MPU.Pasos, MPU.Nivel_actividad)
	return err
}

func (r *MYSQLRepository) DeleteByID(id int) error {
	query := "DELETE FROM mpu WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("no se pudo eliminar el sensor MPU, verifique el id o la sintaxis sql:", err)
		return err
	}
	return nil
}

func (r *MYSQLRepository) UpdateByID(id int, MPU domain.Mpu) error {
	query := "UPDATE mpu SET aceleracion_x=?, aceleracion_y=?, aceleracion_z=?, pasos=?, nivel_actividad=? WHERE id = ?"
	result, err := r.db.Exec(query, MPU.Aceleracion_x, MPU.Aceleracion_y,MPU.Aceleracion_z,MPU.Pasos, MPU.Nivel_actividad, id)

	if err != nil {
		log.Println("no se pudo actualizar el dato, verifique la sintaxis o los datos:", err)
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("error: no se actualizó ningún dato")
	}

	return nil
}

func (r *MYSQLRepository) FindAll() ([]domain.Mpu, error) {
	query := "SELECT aceleracion_x, aceleracion_y, aceleracion_z, pasos, nivel_actividad FROM mpu"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var MPU domain.Mpu
		err := rows.Scan(&MPU.Aceleracion_x, &MPU.Aceleracion_y, &MPU.Aceleracion_z, &MPU.Pasos, &MPU.Nivel_actividad)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, MPU)
	}

	return MPUs, rows.Err()
}

func (r *MYSQLRepository) FindByID(id int) ([]domain.Mpu, error) {
	query := "SELECT * FROM mpu WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MPUs []domain.Mpu
	for rows.Next() {
		var MPU domain.Mpu
		err := rows.Scan(&MPU.Aceleracion_x, &MPU.Aceleracion_y, &MPU.Aceleracion_z, &MPU.Pasos, &MPU.Nivel_actividad)
		if err != nil {
			return nil, err
		}
		MPUs = append(MPUs, MPU)
	}

	return MPUs, rows.Err()
}
