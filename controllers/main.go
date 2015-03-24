package controllers

import (
   "net/http"
   "github.com/golang/glog"

   "github.com/tzjin/snak-attak/system"
   "github.com/tzjin/snak-attak/controllers"
   "github.com/tzjin/snak-attak/models"
)


func hello(c web.C, w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!\n", c.URLParams["name"])
}

func GET_data(c web.C, w http.ResponseWriter, r *http.Request) {
   // send data as json
}

func INC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and incr
}

func DEC_counter(c web.C, w http.ResponseWriter, r *http.Request) {
   // access database and decr
}