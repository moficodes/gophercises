package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "gophercise_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	err = resetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Ping())
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	must(err)
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	must(err)
	return nil
}

/*
pkg: github.com/moficodes/gophercises/phone
BenchmarkNormalize-8   	 1000000	      1209 ns/op	    1024 B/op	      16 allocs/op
*/
func normalize(phone string) string {
	var buf bytes.Buffer
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

/*
pkg: github.com/moficodes/gophercises/phone
BenchmarkNormalize-8   	   30000	     41238 ns/op	  305712 B/op	     198 allocs/op
*/
// func normalize(phone string) string {
// 	re := regexp.MustCompile("\\D")
// 	return re.ReplaceAllString(phone, "")
// }

func must(err error) {
	if err != nil {
		panic(err)
	}
}
