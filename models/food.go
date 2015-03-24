package models

import (
   "fmt"
   "log"
   "database/sql"

   _ "github.com/lib/pq"
)

const (
   DB_USER     = "user"
   DB_PASSWORD = "password"
   DB_NAME     = "mydb"
)

func main() {
   // user and password unncessary?
   dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      DB_USER, DB_PASSWORD, DB_NAME)
   db, err := sql.Open("postgres", dbinfo)
   checkErr(err)
   defer db.Close()

   // do stuff
}

func checkErr(err error, msg string) {
   if err != nil {
      log.Panicln(msg, err)
   }
}
