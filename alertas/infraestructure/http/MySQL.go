package infraestructure

import (
	"API-VITALVEST/alertas/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type MYSQLRepository struct {
	db *sql.DB
}

func NewMYSQLRepository(db *sql.DB) *MYSQLRepository {
	return &MYSQLRepository{db: db}
}

func (r *MYSQLRepository) Save(alerta domain.Alerta) error {
	query := `INSERT INTO alertas (nombre_del_sensor, fecha, cantidad_de_veces_enviado)
	          VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, alerta.NombreDelSensor, alerta.Fecha, alerta.CantidadDeVecesEnviado)
	if err != nil {
		log.Println("Error al guardar la alerta:", err)
	}
	return err
}

func (r *MYSQLRepository) DeleteByID(id int) error {
	query := "DELETE FROM alertas WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la alerta:", err)
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no se eliminó ninguna alerta con id %d", id)
	}
	return nil
}

func (r *MYSQLRepository) UpdateByID(id int, alerta domain.Alerta) error {
	query := `UPDATE alertas 
	          SET nombre_del_sensor = ?, fecha = ?, cantidad_de_veces_enviado = ?
	          WHERE id = ?`
	result, err := r.db.Exec(query, alerta.NombreDelSensor, alerta.Fecha, alerta.CantidadDeVecesEnviado, id)
	if err != nil {
		log.Println("Error al actualizar la alerta:", err)
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no se actualizó ninguna alerta")
	}
	return nil
}

func (r *MYSQLRepository) FindAll() ([]domain.Alerta, error) {
	query := "SELECT id, nombre_del_sensor, fecha, cantidad_de_veces_enviado FROM alertas"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error al obtener alertas:", err)
		return nil, err
	}
	defer rows.Close()

	var alertas []domain.Alerta
	for rows.Next() {
		var alerta domain.Alerta
		err := rows.Scan(&alerta.ID, &alerta.NombreDelSensor, &alerta.Fecha, &alerta.CantidadDeVecesEnviado)
		if err != nil {
			return nil, err
		}
		alertas = append(alertas, alerta)
	}
	return alertas, rows.Err()
}

func (r *MYSQLRepository) FindByID(id int) ([]domain.Alerta, error) {
	query := "SELECT id, nombre_del_sensor, fecha, cantidad_de_veces_enviado FROM alertas WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alertas []domain.Alerta
	for rows.Next() {
		var alerta domain.Alerta
		err := rows.Scan(&alerta.NombreDelSensor, &alerta.Fecha, &alerta.CantidadDeVecesEnviado)
		if err != nil {
			return nil, err
		}
		alertas = append(alertas, alerta)
	}
	return alertas, rows.Err()
}
