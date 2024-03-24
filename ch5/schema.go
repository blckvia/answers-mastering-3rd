// rewrite utility using jackc/pgx
package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}

	host := arguments[1]
	p := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	database := arguments[5]

	// Port number SHOULD BE an integer
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}

	// connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, database)

	connConfig, err := pgx.ParseConfig(conn)
	if err != nil {
		fmt.Println("ParseConfig:", err)
		return
	}

	db, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Println("ConnectConfig:", err)
		return
	}
	defer db.Close(context.Background())

	// Get all databases
	rows, err := db.Query(context.Background(), `SELECT "datname" FROM "pg_database"
	WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		defer rows.Close()
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name)
	}

	// Get all tables from __current__ database
	query := `SELECT table_name FROM information_schema.tables WHERE 
		table_schema = 'public' ORDER BY table_name`
	rows, err = db.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	// This is how you process the rows that are returned from SELECT
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
