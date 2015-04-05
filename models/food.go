package models

import (
   "database/sql"
   "fmt"
   "log"

   "code.google.com/p/go.crypto/bcrypt"
   "github.com/go-gorp/gorp"
   _ "github.com/go-sql-driver/mysql"
   "github.com/golang/glog"
)

type Food struct {
   FoodId      int64    `fid`
   FoodName    string   //`fname`
   Hall        string   //`hall`
   Votes       int32    //`votes`
   Date        string   //`date`
   Meal        string   //`meal`
   Comments    []string //`comments`
}

func InsertFood(dbMap *gorp.DbMap, food *Food) error {
   return dbMap.Insert(food)
}

func GetDbMap(user, password, hostname, port, database string) *gorp.DbMap {
   // connect to db using standard Go database/sql API
   // use whatever database/sql driver you wish
   //TODO: Get user, password and database from config.
   db, err := sql.Open("mysql", fmt.Sprint(user, ":", password, "@(", hostname, ":", port, ")/", database, "?charset=utf8mb4"))
   checkErr(err, "sql.Open failed")

   // construct a gorp DbMap
   dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

   // add a table, setting the table name to 'Food' and
   // specifying that the FoodId property is an auto incrementing PK
   dbMap.AddTableWithName(Food{}, "Foods").SetKeys(true, "FoodId")

   // create the table. in a production system you'd generally
   // use a migration tool, or create the tables via scripts
   err = dbMap.CreateTablesIfNotExists()
   checkErr(err, "Create tables failed")

   return dbMap
}

func checkErr(err error, msg string) {
   if err != nil {
      log.Fatalln(msg, err)
   }
}
