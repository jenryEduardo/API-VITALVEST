package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/session/domain"
	"fmt"
)

type MYSQLRepository struct{
	conn *core.Conn_MySQL
}

func NewMysqlRepo()*MYSQLRepository{
	conn := core.GetDBPool()
	return &MYSQLRepository{conn: conn}
}

func ( r *MYSQLRepository) Save(data domain.Session)error{
	query := "INSERT INTO registro_general VALUES (?,?,?)"
	_,err:= r.conn.DB.Exec(query,&data.UserID,&data.FechaInicio,&data.FechasFin)

	if err!=nil{
		fmt.Print("error verifique su sintaxis sql")
	}

	return err

}
func (r *MYSQLRepository) GetAll() ([]domain.DataSession, error) {
	query := `
	SELECT 
    rg.id AS session_id,
    rg.idUser,
    rg.fecha_inicio,
    rg.fecha_fin,

    bme.temperatura_ambiente,
    bme.humedad_relativa,

    gsr.conductancia,
    gsr.estado_hidratacion,

    mlx.temperatura_corporal,

    mpu.aceleracion_x,
    mpu.aceleracion_y,
    mpu.aceleracion_z,
    mpu.pasos,
    mpu.nivel_actividad

FROM registro_general rg
LEFT JOIN bme ON bme.fecha BETWEEN rg.fecha_inicio AND rg.fecha_fin
LEFT JOIN gsr ON gsr.fecha BETWEEN rg.fecha_inicio AND rg.fecha_fin
LEFT JOIN mlx ON mlx.fecha BETWEEN rg.fecha_inicio AND rg.fecha_fin
LEFT JOIN mpu ON mpu.fecha BETWEEN rg.fecha_inicio AND rg.fecha_fin
WHERE rg.idUser = 1 -- si quieres filtrar por usuario
ORDER BY rg.id`

	rows, err := r.conn.DB.Query(query)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta SQL:", err)
		return nil, err
	}
	defer rows.Close()

	var data []domain.DataSession

	for rows.Next() {
		var session domain.DataSession
		err := rows.Scan(
			&session.SessionID,
			&session.UserID,
			&session.FechaInicio,
			&session.FechaFin,
			&session.TempAmbiente,
			&session.HumedadRelativa,
			&session.Conductancia,
			&session.EstadoHidratacion,
			&session.TempCorporal,
			&session.AceleracionX,
			&session.AceleracionY,
			&session.AceleracionZ,
			&session.Pasos,
			&session.NivelActividad,
		)
		if err != nil {
			fmt.Println("Error al escanear fila:", err)
			continue
		}
		data = append(data, session)
	}

	return data, nil
}
