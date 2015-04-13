package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/context"

	"github.com/tzjin/snak-attak/controllers"
	"github.com/tzjin/snak-attak/system"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

func main() {
	filename := flag.String("config", "config.toml", "Path to configuration file")

	flag.Parse()
	defer glog.Flush()

	var application = &system.Application{}

	application.Init(filename)
	application.LoadTemplates()

	// Setup static files
	static := web.New()
	publicPath := application.Config.Get("general.public_path").(string)
	static.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(publicPath))))

	http.Handle("/assets/", static)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)
	// goji.Use(application.ApplyDbMap)
	goji.Use(application.ApplyAuth)
	goji.Use(application.ApplyIsXhr)
	goji.Use(application.ApplyCsrfProtection)
	goji.Use(context.ClearHandler)

	controller := &controllers.MainController{}
	apicontroller := &controllers.ApiController{}

	// Couple of files - in the real world you would use nginx to serve them.
	goji.Get("/robots.txt", http.FileServer(http.Dir(publicPath)))
	goji.Get("/favicon.ico", http.FileServer(http.Dir(publicPath+"/images")))

	// test code
	// goji.Get("/hello/:name", controller.hello)

	// Home page
	goji.Get("/", application.Route(controller, "Index"))

	// Sign In routes
	goji.Get("/signin", application.Route(controller, "SignIn"))
	goji.Post("/signin", application.Route(controller, "SignInPost"))

	// Sign Up routes
	goji.Get("/signup", application.Route(controller, "SignUp"))
	goji.Post("/signup", application.Route(controller, "SignUpPost"))

	// KTHXBYE
	goji.Get("/logout", application.Route(controller, "Logout"))

	// handlers for /api/* calls
	goji.Get("/api/get/", application.Route(apicontroller, "GET_data"))
	//goji.Get("/api/get/", apicontroller.GET_data)
	// goji.Post("/api/inc/:food", controller.INC_counter)
	// goji.Post("/api/dec/:food", controller.DEC_counter)

	graceful.PostHook(func() {
		application.Close()
	})

	goji.Serve()
}
