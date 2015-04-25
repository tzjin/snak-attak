package controllers

import (
	"fmt"
	"net/http"
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
	
	dbMap := controller.GetDbMap(c)

	var meal string

	// meal
	if time.Now().Hour() < 14 {
		meal = "l"
	} else {
		meal = "d"
	}
	
	msg := models.GetMealData(dbMap, meal)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	fmt.Fprintf(w, "%s\n", msg)
}

func (controller *ApiController) INC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
	// access database and incr
	dbMap := controller.GetDbMap(c)

	name := c.URLParams["food"]

	models.VoteByName(dbMap, name, true)
	fmt.Fprintf(w, "Bingo\n")
}

func (controller *ApiController) DEC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
	// access database and decr
	dbMap := controller.GetDbMap(c)

	name := c.URLParams["food"]

	models.VoteByName(dbMap, name, false)
	fmt.Fprintf(w, "Bingo\n")
}
