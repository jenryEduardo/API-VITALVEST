package infraestructure

import (
	"API-VITALVEST/core"
	"API-VITALVEST/login/domain"
	"fmt"
	 "golang.org/x/crypto/bcrypt"
)

type Mysql_conn struct {
	conn *core.Conn_MYSQL
}

func NewMysqlRepo()*Mysql_conn{
	conn := core.GetDBpool()
	return &Mysql_conn{conn: conn}
}

func (r *Mysql_conn) Login_app(login domain.Login)([]domain.Login,error) {

	query :="SELECT username,passwords FROM users WHERE username = ?"

	rows,err:=r.conn.DB.Query(query,&login.UserName)

	if err!=nil{
		fmt.Println("error",err)
		return nil, err
	}

	defer rows.Close()

	var data []domain.Login

	for rows.Next(){
		var Datos domain.Login
		err:=rows.Scan(&Datos.UserName,&Datos.Password)
		if err!=nil{
			fmt.Println("error no se encontro nada")
		}

		if  bcrypt.CompareHashAndPassword([]byte(Datos.Password),[]byte(login.Password))!=nil{
			fmt.Print("error verifique su contraseña")
		}


	data = append(data, Datos)	
		
	}

	return  data,nil
}
