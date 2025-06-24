package core

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Conn_MYSQL struct{
	DB *sql.DB
	Err string
}

func GetDBpool()*Conn_MYSQL{

	dbUser:=""
	dbPass:=""
	dbHost:=""
	dbSchema:=""

	dsm:= fmt.Sprint("%s:%s@tcp(%s:3306)/%s",dbUser,dbPass,dbHost,dbSchema)

	db,err:=sql.Open("mysql",dsm)
	if err!=nil{
		log.Fatal("error al abrir la BD",err)
	}
	db.SetMaxOpenConns(10) 

	if err:=db.Ping();err!=nil{
		log.Fatal("error ",err)
	}

	return &Conn_MYSQL{DB:db}
}


