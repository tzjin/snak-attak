package main

import ( 
   "fmt"
   "net/http"
   
   "github.com/zenazn/goji"
   "github.com/zenazn/goji/web"

   "github.com/tzjin/snak-attak/system"
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

func main() {
   // test code
   goji.Get("/hello/:name", hello)

   // static file serve
   goji.Get("/*", http.FileServer(http.Dir("public")))

   // handlers for /api/* calls
   goji.Get("/api/get/:food", GET_data)
   goji.Post("/api/inc/:food", INC_counter)
   goji.Post("/api/dec/:food", DEC_counter)
   
   goji.Serve()
}
