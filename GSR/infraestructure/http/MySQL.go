package infraestructure

import (
	"API-VITALVEST/GSR/domain"
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

func (r *MYSQLRepository) Save(gsr domain.Gsr) error {
	query := "INSERT INTO gsr (conductancia, estado_hidratacion) VALUES (?, ?)"
	_, err := r.db.Exec(query, gsr.Conductancia, gsr.Estado_hidratacion)
	return err
}

func (r *MYSQLRepository) DeleteByID(id int) error {
	query := "DELETE FROM gsr WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("no se pudo eliminar el sensor GSR, verifique el id o la sintaxis sql:", err)
		return err
	}
	return nil
}

func (r *MYSQLRepository) UpdateByID(id int, gsr domain.Gsr) error {
	query := "UPDATE gsr SET conductancia=?, estado_hidratacion=? WHERE id = ?"
	result, err := r.db.Exec(query, gsr.Conductancia, gsr.Estado_hidratacion, id)

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

func (r *MYSQLRepository) FindAll() ([]domain.Gsr, error) {
	query := "SELECT conductancia, estado_hidratacion FROM gsr"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gsrs []domain.Gsr
	for rows.Next() {
		var gsr domain.Gsr
		err := rows.Scan(&gsr.Conductancia, &gsr.Estado_hidratacion)
		if err != nil {
			return nil, err
		}
		gsrs = append(gsrs, gsr)
	}

	return gsrs, rows.Err()
}

func (r *MYSQLRepository) FindByID(id int) ([]domain.Gsr, error) {
	query := "SELECT conductancia, estado_hidratacion FROM gsr WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gsrs []domain.Gsr
	for rows.Next() {
		var gsr domain.Gsr
		err := rows.Scan(&gsr.Conductancia, &gsr.Estado_hidratacion)
		if err != nil {
			return nil, err
		}
		gsrs = append(gsrs, gsr)
	}

	return gsrs, rows.Err()
}
