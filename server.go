package main

import ( 
   "fmt"
   "net/http"
   
   "github.com/zenazn/goji"
   "github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!\n", c.URLParams["name"])
}

func main() {
   goji.Get("/hello/:name", hello)
   goji.Get("/*", http.FileServer(http.Dir("public")))
   // create handlers for /api/* calls
   
   goji.Serve()
}
