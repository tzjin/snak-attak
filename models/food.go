package models

import (
   "database/sql"
   "fmt"
   "log"
   
   _ "github.com/lib/pq"
)

const (
   DB_USER     = "user"
   DB_PASSWORD = "password"
   DB_NAME     = "test"
)

func main() {
   dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      DB_USER, DB_PASSWORD, DB_NAME)
   db, err := sql.Open("postgres", dbinfo)
   checkErr(err)
   defer db.Close()
}
