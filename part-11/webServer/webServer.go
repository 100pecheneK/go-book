package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "%s", Body)
	fmt.Fprintf(w, "%s\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName, lastname string
		err = rows.Scan(&id, &firstName, &lastname)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			return
		}
		fmt.Fprintf(w, "%d, %s, %s\n", id, firstName, lastname)
	}
	err = rows.Err()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
}

func getPort() string {
	PORT := ":8001"

	args := os.Args
	if len(args) != 1 {
		PORT = ":" + args[1]
	}
	return PORT
}

func main() {
	PORT := getPort()
	fmt.Println("Using port: ", PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/getData", getDataHandler)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
