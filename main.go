package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=password")

	if err != nil {
		log.Fatal("Unable to connect", err)
	}

	defer conn.Close()

	log.Println("Connected to database")

	// test my connection

	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database")
	}

	log.Println("Pinged the database")

	//get rows from table
	if err = getAllrows(conn); err != nil {
		log.Fatal(err)
	}

	// insert a row
	query := `insert into user_table (name, email) values($1, $2)`
	_, err = conn.Exec(query, "Raihan", "raihan@example.com")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row")

	// get rows from table again
	if err = getAllrows(conn); err != nil {
		log.Fatal(err)
	}

	// update row from table
	stmt := `update user_table set name = $1 where name = $2`

	_, err = conn.Exec(stmt, "Abu Raihan", "Raihan")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated the rows")

	// one row by id
	query = `select id, name, email from user_table where name=$1`

	row := conn.QueryRow(query, "Abu Raihan")
	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
	log.Println("=============")

	// delete one or many row
	query = `delete from user_table where name=$1`
	_, err = conn.Exec(query, "Abu Raihan")
	if err != nil {
		log.Println(err)
	}

	log.Println("Deleted the user")

	// get rows again
	if err = getAllrows(conn); err != nil {
		log.Fatal(err)
	}

}

func getAllrows(conn *sql.DB) error {
	rows, err := conn.Query("select id, name, email from user_table")
	if err != nil {

		log.Fatal("error no", err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal("error no", err)
			return err
		}

		fmt.Println("Record is ", user)
	}

	if err = rows.Err(); err != nil {

		return err
	}

	fmt.Println("========================")

	return nil
}
