package infraestructure

import (
	"API-VITALVEST/MLX/domain"
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

func (r *MYSQLRepository) Save(Mlx domain.Mlx) error {
	query := "INSERT INTO mlx (temperatura_ambiente, temperatura_objeto) VALUES (?, ?)"
	_, err := r.db.Exec(query, Mlx.TemperaturaAmbiente, Mlx.TemperaturaObjeto)
	return err
}

func (r *MYSQLRepository) DeleteByID(id int) error {
	query := "DELETE FROM mlx WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("no se pudo eliminar el sensor Mlx, verifique el id o la sintaxis sql:", err)
		return err
	}
	return nil
}

func (r *MYSQLRepository) UpdateByID(id int, Mlx domain.Mlx) error {
	query := "UPDATE mlx SET temperatura_ambiente = ?, temperatura_objeto=? WHERE id = ?"
	result, err := r.db.Exec(query, Mlx.TemperaturaAmbiente,Mlx.TemperaturaObjeto, id)

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

func (r *MYSQLRepository) FindAll() ([]domain.Mlx, error) {
	query := "SELECT temperatura_ambiente, temperatura_objeto FROM mlx"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Mlxs []domain.Mlx
	for rows.Next() {
		var Mlx domain.Mlx
		err := rows.Scan(&Mlx.TemperaturaAmbiente, &Mlx.TemperaturaObjeto)
		if err != nil {
			return nil, err
		}
		Mlxs = append(Mlxs, Mlx)
	}

	return Mlxs, rows.Err()
}

func (r *MYSQLRepository) FindByID(id int) ([]domain.Mlx, error) {
	query := "SELECT temperatura_ambiente, temperatura_objeto FROM mlx WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Mlxs []domain.Mlx
	for rows.Next() {
		var Mlx domain.Mlx
		err := rows.Scan(&Mlx.TemperaturaAmbiente, &Mlx.TemperaturaObjeto)
		if err != nil {
			return nil, err
		}
		Mlxs = append(Mlxs, Mlx)
	}

	return Mlxs, rows.Err()
}
