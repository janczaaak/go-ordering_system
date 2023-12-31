package config

import(
	"database/sql"
)

func ConnectDb() *sql.DB{
	dbDriver := "mysql"
	dbUser := "your_username"
	dbPass := "your_password"
	dbName := "orders"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	return db
}
