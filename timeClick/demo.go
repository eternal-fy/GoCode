package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	defer c.Stop()
	spec := "*/1 * * * * ?" //crom  s m h d M date
	i := 0
	c.AddFunc(spec, func() {
		fmt.Println(i)
		i++
	})
	c.Start()
	select {}
}
