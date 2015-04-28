package main

import (
	"fmt"
	"github.com/robfig/cron"

	"os"
	"os/exec"
)

func main() {
	c := cron.New()

	//c.AddFunc("@every 1day", func() { fmt.Println("Every hour thirty") })
	c.Start()
	//for{

	// Scrape data
	os.Chdir("./helpers")
	a, err := exec.Command("python", "scrape.py").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		// Store data somewhere
		fmt.Printf("\n%s\n")
		/*for _, hall := range a {
		  if length(hall) == 1 {
		    fmt.Print(hall)
		  }
		}*/
	}
	//}
}
