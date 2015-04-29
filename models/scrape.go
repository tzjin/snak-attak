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
	// Roma
	for _, a := range h.Roma.Breakfast {
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Roma.Lunch {
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Roma.Dinner {
		fo := Food{Name: a.Name, Hall: "Rocky/Mathey", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
		f = append(f, fo)
	}

	// Wucox
	for _, a := range h.Wucox.Breakfast {
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Wucox.Lunch {
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Wucox.Dinner {
		fo := Food{Name: a.Name, Hall: "Wu/Wilcox", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
		f = append(f, fo)
	}

	// Whitman
	for _, a := range h.Whitman.Breakfast {
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Whitman.Lunch {
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Whitman.Dinner {
		fo := Food{Name: a.Name, Hall: "Whitman", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
		f = append(f, fo)
	}
	// Forbes
	for _, a := range h.Forbes.Breakfast {
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Forbes.Lunch {
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Forbes.Dinner {
		fo := Food{Name: a.Name, Hall: "Forbes", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
		f = append(f, fo)
	}
	// Grad
	for _, a := range h.Grad.Breakfast {
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.Grad.Lunch {
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.Grad.Dinner {
		fo := Food{Name: a.Name, Hall: "Grad", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
		f = append(f, fo)
	}
	// CJL
	for _, a := range h.CJL.Breakfast {
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: "today", Filters: a.Filt, Meal: "b"}
		f = append(f, fo)
	}
	for _, a := range h.CJL.Lunch {
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: "today", Filters: a.Filt, Meal: "l"}
		f = append(f, fo)
	}
	for _, a := range h.CJL.Dinner {
		fo := Food{Name: a.Name, Hall: "CJL", Votes: 0, Date: "today", Filters: a.Filt, Meal: "d"}
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
