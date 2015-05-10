package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	// Get foods that match name + today's date
	obj, err := dbMap.Get(Food{}, foodid)

	food, ok := obj.(*Food)
	if !ok {
		glog.Warningf("Cannot convert interface: %v", err)
	}

	if up {
		food.Votes++
	} else {
		food.Votes--
	}

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
	Eastern := time.FixedZone("Eastern", -4*3600)
	today := time.Now().UTC().In(Eastern).Format("01-02-2006")
	fmt.Printf("%v\n", time.Now().UTC().In(Eastern))
	query := "SELECT * FROM foods where date='" + today + "' ORDER BY votes DESC "
	_, err := dbMap.Select(&foods, query)

	if err != nil {
		glog.Warningf("Can't get foods by meal: %v", err)
	}

	return
}

// write get comments
func GetDbMap() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err, "postgres.Open failed")

	// construct a gorp DbMap
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'Foods' and
	// specifying that the FoodId property is an auto incrementing PK
	t := dbMap.AddTableWithName(Food{}, "foods").SetKeys(true, "Id")
	t.ColMap("Name").SetMaxSize(100)
	t.ColMap("Meal").SetMaxSize(1)

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
