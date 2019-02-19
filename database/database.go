package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetConnection() *sql.DB {
	driver := "mysql"  //driver do banco a ser usado
	user := "daniel"   //usuário do banco
	password := "1234" //senha do usuario
	banco := "projeto_usuarios"
	db, err := sql.Open(driver, user+":"+password+"@/"+banco+"?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
