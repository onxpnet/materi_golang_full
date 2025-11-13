package db
import (
	"database/sql"
	_ "github.com/lib/pq"
	"esdm/env"
	"fmt"
	"log"
)

var DB *sql.DB

func Init() {
	env.LoadEnvFile(".env")
	connection_string := env.GetEnv("DATABASE_URL", "")

	var err error
	DB, err = sql.Open("postgres", connection_string)
	
	if err != nil {
		panic(err)
	}
	
	if err = DB.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	//defer DB.Close()
}
