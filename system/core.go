package system

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	// "crypto/sha256"

	"github.com/go-gorp/gorp"
	"github.com/golang/glog"
	// "github.com/gorilla/sessions"
	"github.com/robfig/cron"
	"github.com/zenazn/goji/web"
	"sniksnak/models"
)

type CsrfProtection struct {
	Key    string
	Cookie string
	Header string
	Secure bool
}

type Application struct {
	Template *template.Template
	DbMap    *gorp.DbMap
}

func (application *Application) Init() {

	application.DbMap = models.GetDbMap()

	// Setup scheduler + scraper
	// runs a minute after the hour
	c := cron.New()
	c.AddFunc("@midnight", func() {
		if len(models.GetFoodByMeal(application.DbMap, "l")) == 0 {
			models.StoreDailyData(application.DbMap)
		}
	})
	c.Start()
}

func (application *Application) LoadTemplates() error {
	var templates []string

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	err := filepath.Walk("views", fn)

	if err != nil {
		return err
	}

	application.Template = template.Must(template.ParseFiles(templates...))
	return nil
}

func (application *Application) Close() {
	glog.Info("Bye!")
}

func (application *Application) Route(controller interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		c.Env["Content-Type"] = "text/html"

		methodValue := reflect.ValueOf(controller).MethodByName(route)
		methodInterface := methodValue.Interface()
		method := methodInterface.(func(c web.C, r *http.Request) (string, int))

		body, code := method(c, r)

		switch code {
		case http.StatusOK:
			if _, exists := c.Env["Content-Type"]; exists {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			}
			io.WriteString(w, body)
		case http.StatusSeeOther, http.StatusFound:
			http.Redirect(w, r, body, code)
		}
	}
	return fn
}
