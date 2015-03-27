package controllers

import (
   "fmt"
   // "log"
   "net/http"

   "github.com/zenazn/goji/web"

   "github.com/tzjin/snak-attak/system"
   "github.com/tzjin/snak-attak/models"
)

type MainController struct {
   system.Controller
}

func (controller *MainController) hello(c web.C, w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!\n", c.URLParams["name"])
}

func (controller *MainController) GET_data(c web.C, w http.ResponseWriter, r *http.Request) {
   // send data as json
}

func (controller *MainController) INC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and incr
}

func (controller *MainController) DEC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and decr
}