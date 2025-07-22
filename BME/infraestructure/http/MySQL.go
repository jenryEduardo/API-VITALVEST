package infraestructure

import (
	"API-VITALVEST/BME/domain"
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

func (r *MYSQLRepository) Save(BME domain.Bme) error {
	query := "INSERT INTO bme (temperatura_ambiente, presion,humedad) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, BME.Temperatura, BME.Humedad, BME.Presion)
	return err
}

func (r *MYSQLRepository) DeleteByID(id int) error {
	query := "DELETE FROM bme WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("no se pudo eliminar el sensor BME, verifique el id o la sintaxis sql:", err)
		return err
	}
	return nil
}

func (r *MYSQLRepository) UpdateByID(id int, BME domain.Bme) error {
	query := "UPDATE bme SET temperatura=?, humedad=?, presion=?WHERE id = ?"
	result, err := r.db.Exec(query,BME.Temperatura, BME.Humedad, BME.Presion, id)

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

func (r *MYSQLRepository) FindAll() ([]domain.Bme, error) {
	query := "SELECT temperatura, humedad, presion FROM bme"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var BMEs []domain.Bme
	for rows.Next() {
		var BME domain.Bme
		err := rows.Scan(&BME.Temperatura, &BME.Humedad, &BME.Presion)
		if err != nil {
			return nil, err
		}
		BMEs = append(BMEs, BME)
	}

	return BMEs, rows.Err()
}

func (r *MYSQLRepository) FindByID(id int) ([]domain.Bme, error) {
	query := "SELECT temperatura, humedad, presion FROM bme WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var BMEs []domain.Bme
	for rows.Next() {
		var BME domain.Bme
		err := rows.Scan(&BME.Temperatura, &BME.Humedad, &BME.Presion)
		if err != nil {
			return nil, err
		}
		BMEs = append(BMEs, BME)
	}

	return BMEs, rows.Err()
}
