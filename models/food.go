package models

import (
   "database/sql"
   "fmt"
   "log"

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

//Todo: figure out interfaces in go


// func GetFoodByHall(dbMap *gorp.DbMap, hall string) (foods *Food) {
//    err := dbMap.Select(&foods, "SELECT * FROM Foods where Hall = ?", hall)

//    if err != nil {
//       glog.Warningf("Can't get foods by dining hall: %v", err)
//    }
//    return
// }

func GetFoodByID(dbMap *gorp.DbMap, foodid int64) (food *Food) {
   fud, err := dbMap.Get(Food{}, foodid)

   if err != nil {
      glog.Warningf("Can't get foods by id: %v", err)
   }

   food, ok := fud.(*Food)
   if !ok {
      // cannot convert interface
   }

   return 
}

func GetFoodByMeal(dbMap *gorp.DbMap, meal string) (foods []*Food) {
   _, err := dbMap.Select(&foods, "SELECT * FROM Foods where Meal = ?", meal)

   if err != nil {
      glog.Warningf("Can't get foods by meal: %v", err)
   }
   return
}

func GetCommentsForID(dbMap *gorp.DbMap, id int64) (comments []string) {
   food, err := dbMap.Get(Food{}, id)

   if err != nil {
      glog.Warningf("Can't get comments of id: %v", err)
   }

   items, ok := food.(*Food)

   if !ok {
      // cannot convert interface
   }
   comments = items.Comments
   return
}

// func GetVotesForID(dbMap *gorp.DbMap, id int64) (votes int32) {
//    food, err := dbMap.Get(Food{}, id)
//    items, ok := food.(*Food)
//    votes = items.Votes
//    return
// }

func GetFoodDbMap(user, password, hostname, port, database string) *gorp.DbMap {
   // connect to db using standard Go database/sql API
   // use whatever database/sql driver you wish
   //TODO: Get user, password and database from config.
   db, err := sql.Open("mysql", fmt.Sprint(user, ":", password, "@(", hostname, ":", port, ")/", database, "?charset=utf8mb4"))
   checkErr(err, "sql.Open failed")

   // construct a gorp DbMap
   dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

   // add a table, setting the table name to 'Foods' and
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
