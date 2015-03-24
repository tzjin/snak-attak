package models

import (
   "fmt"
   "log"
   "database/sql"

   "github.com/go-gorp/gorp"  // map structs to db
   "github.com/golang/glog"   // nicer log package
   _ "github.com/lib/pq"      // postgres driver
)

type Food struct {
   // lowercase unnecessary
   Id       int64    'id'
   Name     string   'name'
   Hall     string   'hall'
   Rating   int32    'rating'
}

func GetDbMap(usr, pwd, dbname string) *gorp.DbMap{
   // user and password unncessary?
   dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      usr, pwd, dbname)
   db, err := sql.Open("postgres", dbinfo)
   checkErr(err, "Could not open database")
   // defer db.Close()

   dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

   dbMap.AddTableWithName(Food{}, "Food").SetKeys(true, "Id")

   err = dbMap.CreateTablesIfNotExists()
   checkErr(err, "Create Table Failed")

   return dbMap
}

func checkErr(err error, msg string) {
   if err != nil {
      glog.Panicln(msg, err)
   }
}
