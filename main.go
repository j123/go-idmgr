package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	name     string
	email    string
	password string
	sex      string
	created  string
}

func main() {
	// Connect to sample db
	db, err := sql.Open("mysql", "root:password@/my_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User

	err = db.QueryRow(`
		SELECT
			id,
			name,
			email,
			password,
			sex,
			created
		FROM
			users
		WHERE
			name = 'C17'
	`).Scan(&(user.id), &(user.name), &(user.email), &(user.password), &(user.sex), &(user.created))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(user)

	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// columns, err := rows.Columns()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }
	// for rows.Next() {
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	var value string
	// 	for i, col := range values {
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		fmt.Println(columns[i], ": ", value)
	// 	}
	// 	fmt.Println("-----------------------------------")
	// }
}
