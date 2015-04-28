package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	for {
	}
}
