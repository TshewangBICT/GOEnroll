package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// database details
const (
	postgres_host     = "dpg-d86k0v6q1p3s73c3mmg0-a.singapore-postgres.render.com"
	postgress_port    = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "OCj0FJsINrKSotVlHSS1VSHsy7s5A1uM"
	postgres_dbname   = "my_db_77a0"
)

// DB variable to store the address of our database
var Db *sql.DB

func init() {
	// create a connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", postgres_host, postgress_port, postgres_user, postgres_password, postgres_dbname)
	var err error
	// establish the connection to database server using the driver (lib/pq)
	Db, err = sql.Open("postgres", db_info)

	// Handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully established")
	}

}
