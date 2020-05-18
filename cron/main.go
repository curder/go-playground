package main

import (
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main()  {
	InitCron()

	for {
		time.Sleep(time.Second * 3)
	}
}

func InitCron() {
	var (
		c   *cron.Cron
		entryID cron.EntryID
		err error
	)

	c = cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	if entryID, err = c.AddFunc("* * * * *", RunSomething); err != nil {
		panic(err)
	}

	log.Println("current entry id:", entryID)

	go c.Start()

	defer c.Stop()
}

