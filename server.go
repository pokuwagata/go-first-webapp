package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])

	var datasource string
	if os.Getenv("env") == "PROD" {
		datasource = os.Getenv("CLEARDB_USER") +
			 ":" +
			 os.Getenv("CLEARDB_PASSWORD") +
			 "@tcp(" + 
			 os.Getenv("CLEARDB_HOST") +
			 ":3306)/" +
			 os.Getenv("CLEARDB_DBNAME")
	} else {
		datasource = "user:password@(db:3306)/sample_db"
	}

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(writer, "%v", db)

	err = db.Ping()
	if err != nil {
		// do something here
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}