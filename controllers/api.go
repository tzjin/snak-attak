package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zenazn/goji/web"

	"sniksnak/models"
	"sniksnak/system"
)

type ApiController struct {
	system.Controller
}

func (controller *ApiController) hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", c.URLParams["name"])
}

func (controller *ApiController) GET_data(c web.C, w http.ResponseWriter, r *http.Request) {
	//dbMap := controller.GetDbMap(c)

	var meal string

	// meal
	if time.Now().Hour() < 14 {
		meal = "Lunch"
	} else {
		meal = "Dinner"
	}

	msg := models.GetMealData(nil, meal)
	fmt.Fprintf(w, "%s\n", msg)
}

func (controller *ApiController) INC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
	// access database and incr
	dbMap := controller.GetDbMap(c)

	id, err := strconv.ParseInt(c.URLParams["food"], 10, 64)

	if err != nil {

	}

	models.VoteById(dbMap, id, true)
	fmt.Fprintf(w, "Bingo\n")
}

func (controller *ApiController) DEC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
	// access database and decr
	dbMap := controller.GetDbMap(c)

	id, err := strconv.ParseInt(c.URLParams["food"], 10, 64)

	if err != nil {

	}

	models.VoteById(dbMap, id, false)
	fmt.Fprintf(w, "Bingo\n")
}
