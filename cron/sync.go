package main

import (
	"log"
)

func RunSomething() {
	log.Println("task1 start running...")
	go task1()
	log.Println("task2 start running...")
	go task2()
	log.Println("task3 start running...")
	go task3()
}

func task1() {
	log.Println("task1 func running...")
}

func task2() {
	log.Println("task2 func running...")
}

func task3() {
	log.Println("task3 func running...")
}
