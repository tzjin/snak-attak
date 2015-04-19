package models

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"database/sql"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

type Food struct {
	FoodId   int    `FoodIid`
	Name     string `FoodName`
	Hall     string `hall`
	Votes    int    `votes`
	Date     string `date`
	Meal     string `meal`
	Comments []int  `comments`
	// Filters?
}

// useful in future for not sending everthing
type Message struct {
	Id      int64
	Name    string
	Hall    string
	Votes   int64
	Filters []string // includes
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

func GetMealData(dbMap *gorp.DbMap, meal string) string {

	var msg bytes.Buffer
	first := true

	foods := GetFoodByMeal(dbMap, meal)

	// build json message
	msg.WriteString("[")

	for i := 0; i < len(foods); i++ {
		if !first {
			msg.WriteString(", ")
		}

		b, err := json.Marshal(foods[i])

		if err != nil {
			glog.Warningf("Cannot encode json: %v", err)
		}

		msg.WriteString(string(b[:]))
		first = false
	}

	msg.WriteString("]")

	return msg.String()
}

func VoteByName(dbMap *gorp.DbMap, foodname string, up bool) (food *Food) {
	// Today's date
	t := time.Now().Local()
	date := t.Format("2006-01-02")

	// Get foods that match name + today's date
	fuds, err := dbMap.Select(Food{}, "SELECT * FROM Foods where fname = ? and date = ? ", foodname, date)

	food, ok := fuds[0].(*Food)
	if !ok {
		// cannot convert interface
	}

	if up {
		food.Votes++
	} else {
		food.Votes--
	}
	count, err := dbMap.Update(&food)

	if err != nil {
		glog.Warningf("Update votes by ID failed: %v", err)
	}

	if count != 1 {
		glog.Warningf("Too many foods updated: %v", err)
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

/*func GetCommentsForID(dbMap *gorp.DbMap, id int64) (comments []string) {
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
}*/

func GetDbMap() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	//checkErr(err, "postgres.Open failed")
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

	// add a table, setting the table name to 'Foods' and
	// specifying that the FoodId property is an auto incrementing PK
	t := dbMap.AddTableWithName(Food{}, "Foods").SetKeys(true, "FoodId")
	t.ColMap("foodname").SetMaxSize(30)
	t.ColMap("foodname").SetMaxSize(20)
	t.ColMap("meal").SetMaxSize(1)

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
