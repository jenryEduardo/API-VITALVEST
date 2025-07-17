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

func (r *MYSQLRepository)GetAll()([]domain.Session,error){
	query := "SELECT * FROM registro_general"
	rows,err:=r.conn.DB.Query(query)

	if err!=nil{
		fmt.Print("verifique la sintaxis sql")
	}

	defer rows.Close()

	var data []domain.Session

	for rows.Next(){
		var Data domain.Session
		err := rows.Scan(&Data.UserID,&Data.FechaInicio,&Data.FechasFin)

		if err!=nil{
			fmt.Print("error no se obtuvo datos")
		}

		data = append(data, Data)
	}

	return data,err
}