package main

import (
	"fmt"
	cron "github.com/robfig/cron"
)

func main()  {
	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Every hour on the half hour")
		c.Stop()
	})
	c.Run()
}