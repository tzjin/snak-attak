package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	t := time.NewTicker(15 * time.Minute)
	// or just use the usual for { select {} } idiom of receiving from a channel
	for _ = range t.C {
	}
}
