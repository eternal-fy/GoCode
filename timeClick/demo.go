package main

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
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
	Click()
	c.Start()

	select {}
}

func Click() {
	cron := "0/1 * * * * *"
	tk := toolbox.NewTask("mytask", cron, func() error {
		return Say()
	})
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()

	defer toolbox.StopTask()
	select {}
}
func Say() error {
	fmt.Println("hehe")
	return nil
}
