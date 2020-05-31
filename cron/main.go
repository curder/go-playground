package main

import (
	"github.com/robfig/cron/v3"
)

func main()  {
	InitCron()
}

func InitCron() {
	var (
		spec string
		c   *cron.Cron
	)
	
	// c = cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	c = cron.New(cron.WithSeconds())
	
	spec = "*/1 * * * * ?"

	c.AddFunc(spec, RunSomething)

	// 启动计划任务
        go c.Start()

	defer c.Stop()

	select{}
}

