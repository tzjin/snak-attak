package main

import ( 
   "fmt"
   "net/http"
   
   "github.com/zenazn/goji"
   "github.com/zenazn/goji/web"

   "github.com/tzjin/snak-attak/system"
   "github.com/tzjin/snak-attak/controllers"
)

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
