package controllers

import (
   "fmt"
   "encoding/json"
   "net/http"

   "github.com/zenazn/goji/web"
   
   "github.com/tzjin/snak-attak/models"
   "github.com/tzjin/snak-attak/system"
)

type ApiController struct {
	system.Controller
}

type Message struct {
   Id          int64
   Name        string
   Location    string
   Votes       int64
   // filter array?
}

func (controller *ApiController) hello(c web.C, w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!\n", c.URLParams["name"])
}

func (controller *ApiController) GET_data(c web.C, w http.ResponseWriter, r *http.Request) (string, int){
   // send all data as json
   // cmnts := []string{}
   // fud := Food{1234, "Chicken Tenders", "Wilson", 23, "4/9/15", "Lunch", cmnts}

   // b, err := json.Marshal(fud)
   // return string(b[:]), http.StatusOK

   // return
}

func (controller *ApiController) INC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and incr
}

func (controller *ApiController) DEC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and decr
}