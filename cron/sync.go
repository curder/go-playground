package main

import (
	"log"
	"time"
)

// See: https://blog.csdn.net/tianshi418/article/details/104695932

var isTaskFinished = 0

func RunSomething() {
	if isTaskFinished == 0 {
		isTaskFinished = 1
		task1()
		task2()
		task3()
		isTaskFinished = 0
		log.Println("task end")
	}
}

func task1() {
	log.Println("task1 func running...")
	time.Sleep(time.Second * 5)
}

func task2() {
	log.Println("task2 func running...")
}

func task3() {
	log.Println("task3 func running...")
	time.Sleep(time.Second * 6)
}