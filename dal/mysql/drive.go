package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitMysql() {
	DB, _ = sql.Open("mysql", "ais_r:TrBLnIE2DvDFDHS6@tcp(123.206.67.38:3306)/ais")
	if err := DB.Ping(); err != nil{
		log.Println("opon database fail")
		return
	}
	log.Println("connnect database success")
}