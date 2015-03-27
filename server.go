package main

import ( 
   "net/http"
   
   "github.com/zenazn/goji"

   "github.com/tzjin/snak-attak/system"
   "github.com/tzjin/snak-attak/controllers"
)

func main() {

   app = &controllers.MainController{}
   // test code
   goji.Get("/hello/:name", app.hello)

   // static file serve
   goji.Get("/*", http.FileServer(http.Dir("public")))

   // handlers for /api/* calls
   goji.Get("/api/get/:food", app.GET_data)
   goji.Post("/api/inc/:food", app.INC_counter)
   goji.Post("/api/dec/:food", app.DEC_counter)
   
   goji.Serve()
}
