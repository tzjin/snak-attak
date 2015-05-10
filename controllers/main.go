package controllers

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"html/template"

	"sniksnak/helpers"
	"sniksnak/models"
	"sniksnak/system"
)

type MainController struct {
	system.Controller
}

// Home page route
func (controller *MainController) Index(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)

	dbMap := controller.GetDbMap(c)
	widgets := helpers.Parse(t, "home", nil)

	// With that kind of flags template can "figure out" what route is being rendered
	c.Env["IsIndex"] = true
	c.Env["Foods"] = models.GetFoodByMeal(dbMap, "l")

	c.Env["Title"] = "SnikSnak"
	c.Env["Content"] = template.HTML(widgets)

	return helpers.Parse(t, "main", c.Env), http.StatusOK
}
