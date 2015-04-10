package models

import (
   "fmt"
   "log"
   "bytes"
   "encoding/json"

   "database/sql"

   "github.com/go-gorp/gorp"
   _ "github.com/lib/pq"
   "github.com/golang/glog"
)

type Food struct {
   Id          int64    `fid`
   Name        string   `FoodName`
   Hall        string   //`hall`
   Votes       int32    //`votes`
   Date        string   //`date`
   Meal        string   //`meal`
   Comments    []string //`comments`
   // Filters?
}

// useful in future for not sending everthing
type Message struct {
   Id          int64
   Name        string
   Hall        string
   Votes       int64
   Filters     []string // includes 
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

// func GetMealData(dbMap *gorp.DbMap, meal string) string {
func GetMealData() string {
   filt := []string{"Vegan", "Victorfood"}
   tndrs := Food{1234, "Chicken Tenders", "Wilson", 23, ,"December 31, 1999", "Dinner", filt}
   foods := []Food{tndrs, tndrs, tndrs}

   var msg bytes.Buffer
   first := true

   // foods := GetFoodByMeal(dbMap, meal)

   // build json message
   msg.WriteString( "[" );

   for i := 0; i < len(foods); i++ {
      if !first {
         msg.WriteString( ", " )
      }

      b, err := json.Marshal(foods[i])

      if err != nil {
         glog.Warningf("Cannot encode json: %v", err)
      }

      msg.WriteString( string(b[:]) )
      first = false
   }

   msg.WriteString( "]" );

   return msg.String()
}

func VoteById(dbMap *gorp.DbMap, foodid, vote int32) (food *Food) {
   fud, err := dbMap.Get(Food{}, foodid)

   if err != nil {
      glog.Warningf("Can't get foods by id: %v", err)
   }

   food, ok := fud.(*Food)
   if !ok {
      // cannot convert interface
   }

   food.Votes += vote
   count, err := dbMap.update(&food)

   if err != nil {
      glog.Warningf("Update votes by ID failed: %v", err)
   }

   return 
}

func GetFoodByMeal(dbMap *gorp.DbMap, meal string) (foods []*Food) {
   // meal of today?
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

func GetDbMap(user, password, hostname, port, database string) *gorp.DbMap {
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
