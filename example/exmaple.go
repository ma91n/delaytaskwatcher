package main

import (
	"fmt"
	"github.com/ma91n/delaytaskwatcher"
	"time"
)

func main() {

	go delaytaskwatcher.Watch(delaytaskwatcher.Options{
		Limit: 10 *time.Second,
	})

	heavyTask()

}

func heavyTask() {
	fmt.Println("start heavytask")
	heavyTaskMain()
	fmt.Println("end heavy")
}

func heavyTaskMain() {
	time.Sleep(30 * time.Second)
}
