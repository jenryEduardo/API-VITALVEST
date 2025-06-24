package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/users/domain"
	"log"
	
) 

type MYSQLRepository struct{
	conn *core.Conn_MYSQL
}

func NewMysqlRepo()*MYSQLRepository{
	conn := core.GetDBpool()
	return &MYSQLRepository{conn: conn}
}

func (r *MYSQLRepository)Save(user domain.User)error{

	query:="INSERT INTO users(name,age,fech_nac) VALUES(?,?,?)"
	_,err:=r.conn.DB.Exec(query,&user.Name,&user.Age,&user.Fech_nac)
	if err!=nil{
		return err
	}

	return err

}

func (r *MYSQLRepository)Delete(id int)error{

	query := "DELETE FROM users WHERE id_user = ?"
	_,err:=r.conn.DB.Exec(query,id)
		if err!=nil{
			log.Fatal("no se pudo eliminar al usuario verifique el id o la siintaxis sql")
		}

return nil
}

func (r *MYSQLRepository)Update(user domain.User,id int)error{
	query := "UPDATE users(name,age,fech_nac) values(?,?,?) WHERE id_user = ?"

	_,err :=r.conn.DB.Exec(query,&user.Name,&user.Age,&user.Fech_nac,&id)

		if err!=nil{

		}

		return err
}