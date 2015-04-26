package models

import (
	"bytes"
	"encoding/json"
	"log"
	// "time"
	"fmt"

	"database/sql"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

type Food struct {
	Id       int    //`db:"foodid"`
	Name     string //`db:"fname"`
	Hall     string //`db:"hall`
	Votes    int    //`db:"votes`
	Date     string //`db:"date`
	Meal     string //`db:"meal`
	Filters  string //'db:"filters
	Comments string //`comments`
	// Filters?
}

func InsertFood(dbMap *gorp.DbMap, food *Food) error {
	return dbMap.Insert(food)
}

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

func VoteById(dbMap *gorp.DbMap, foodid int64, up bool) (food *Food) {
	// Today's date
	// t := time.Now().Local()
	// date := t.Format("2006-01-02")

	// Get foods that match name + today's date
	// fuds, err := dbMap.Select(Food{}, "SELECT * FROM Foods where fname = $1 and date = $2 ", foodid, date)
	obj, err := dbMap.Get(Food{}, foodid)

	food, ok := obj.(*Food)
	if !ok {
		// cannot convert interface
	}

	if up {
		food.Votes++
	} else {
		food.Votes--
	}

	fmt.Printf("%d\n", food.Votes)

	count, err := dbMap.Update(food)

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
	_, err := dbMap.Select(&foods, "SELECT * FROM foods ORDER BY votes DESC")

	if err != nil {
		glog.Warningf("Can't get foods by meal: %v", err)
	}

	return
}

// write get comments

func GetDbMap() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	//checkErr(err, "postgres.Open failed")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'Foods' and
	// specifying that the FoodId property is an auto incrementing PK
	t := dbMap.AddTableWithName(Food{}, "foods").SetKeys(true, "Id")
	t.ColMap("Name").SetMaxSize(30)
	// t.ColMap("foodname").SetMaxSize(20)
	t.ColMap("Meal").SetMaxSize(1)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbMap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// food := Food {Name: "Chicken Tenders", Hall: "Wu/Wilcox", Votes: 32, Date: "today", Filters: "Victorfood", Meal: "d" }
	// err = dbMap.Insert(&food)

	// var foods []Food
	// _, err = dbMap.Select(&foods, "SELECT * FROM foods")
	// fmt.Printf("%d\n",len(foods))
	// for x, p := range foods {
	// 	fmt.Printf("    %d: %v\n", x, p)
	// }

	return dbMap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
