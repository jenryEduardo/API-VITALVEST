package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/sesion/domain"
	"fmt"
)

type MYSQLRepository struct {
	conn *core.Conn_MySQL
}


func NewMysqlRepo()*MYSQLRepository{
	conn := core.GetDBPool()
	return &MYSQLRepository{conn: conn}
}


func(r *MYSQLRepository)Save(sesion domain.Sesion)error{
	query:="INSERT INTO sesions(id_usuario,fech_creacion) VALUES(?,?)"
	_,err:= r.conn.DB.Exec(query,&sesion.Id_usuario,&sesion.Fech_creacion)

		if err!=nil{
			fmt.Println("error no se pudo realizar la solicitud a la BD")
		}

	return err
}


func (r *MYSQLRepository)Update(session domain.Sesion,id int)error{
	query:="UPDATE session SET id_usuario=?,fech_creacion=? WHERE id_sesion=?"
	_,err:=r.conn.DB.Exec(query,&session.Id_usuario,&session.Fech_creacion,id)

		if err!=nil{
			fmt.Println("no se pudo realizar la solicitud a la BD verifique los campos o la ruta",err)
		}

	return err
}


func (r *MYSQLRepository)Delete(id int)error{
	query:="DELETE FROM session WHERE id_sesion=?"
	_,err:=r.conn.DB.Exec(query,id)

		if err!=nil{
			fmt.Println("ocurrio un error revice la sintaxis o el id no existe")
		}

		return err
}


func (r *MYSQLRepository)GetAll()([]domain.Sesion,error){

	var Data []domain.Sesion

	query:="SELECT id_usuario,fech_creacion FROM session"
	rows,err:=r.conn.DB.Query(query)
		if err!=nil{
			fmt.Println("revice la query de la solicitud")
		}

	defer rows.Close()

	for rows.Next(){
		var data domain.Sesion
		rows.Scan(&data.Id_sesion,&data.Id_usuario,&data.Fech_creacion)

			if err!=nil{
				fmt.Print("ocurruo un error revice los tipos de datos que recibe")
			}

			Data = append(Data, data)

			if err:=rows.Err();err!=nil{
			fmt.Print("no se puedo  obtener la lista completa")
			}

		
	}

		return Data,err
}


func (r *MYSQLRepository)GetById(id int)([]domain.Sesion,error){
	query:="SELECT id_sesion,id_usuario,fech_creacion FROM session WHERE id_sesion=?"
	rows,err:=r.conn.DB.Query(query,id)

		if err!=nil{
			fmt.Print("no se pudo realziar la solicitud a la BD")
		}

	defer rows.Close()

	var Data []domain.Sesion

	for rows.Next(){
		var data domain.Sesion
		rows.Scan(&data.Id_sesion,&data.Id_usuario,&data.Fech_creacion)

		if err!=nil{
			fmt.Print("ocurrip un error al escanear los datos verificar tipos de datos de la clase session")
		}
	
		Data = append(Data, data)

	}

	return Data,err
}