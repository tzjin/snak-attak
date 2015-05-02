package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-gorp/gorp"
	"os"
	"os/exec"
)

type SFood struct {
	Name string
	Filt []string
}

type Meals struct {
	Breakfast []SFood `json:"Breakfast"`
	Lunch     []SFood `json:"Lunch"`
	Dinner    []SFood `json:"Dinner"`
}

type Halls struct {
	Roma    Meals `json:"roma"`
	Wucox   Meals `json:"wucox"`
	Whitman Meals `json:"whitman"`
	Forbes  Meals `json:"forbes"`
	Grad    Meals `json:"grad"`
	CJL     Meals `json:"cjl"`
}

func toFood(h Halls) []Food {
	f := []Food{}
	today := time.Now().Format("01-02-2006")
	// Roma
	for _, a := range h.Roma.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Roma.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Roma.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}

	// Wucox
	for _, a := range h.Wucox.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Wucox.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Wucox.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}

	// Whitman
	for _, a := range h.Whitman.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Whitman.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Whitman.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}
	// Forbes
	for _, a := range h.Forbes.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Forbes.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Forbes.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}
	// Grad
	for _, a := range h.Grad.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Grad.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Grad.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}
	// CJL
	for _, a := range h.CJL.Breakfast {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: today, Filters: filter, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.CJL.Lunch {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: today, Filters: filter, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.CJL.Dinner {
		filter := ""
		for i, fil := range a.Filt {
			if i > 0 {
				filter += ", "
			}
			filter += fil
		}
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: today, Filters: filter, Meal: "d"}
		f = append(f, fo)
	}
	return f
}

func StoreDailyData(dbMap *gorp.DbMap) {
	// Scrape data
	os.Chdir("./helpers")
	a, err := exec.Command("python", "scrape.py").Output()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		// Unmarshall Json
		var h Halls
		err := json.Unmarshal([]byte(a), &h)
		if err != nil {
			fmt.Println(err)
		}

		// Store data into database
		allFoods := toFood(h)
		for _, f := range allFoods {
			err = dbMap.Insert(&f)
			if err != nil {
				fmt.Println("Error: %v\n", err)
			}

		}
	}
}
