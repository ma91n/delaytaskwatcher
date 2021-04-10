package main

import (
	"fmt"
	"github.com/ma91n/delaytaskwatcher"
	"time"
)

func main() {

	go delaytaskwatcher.Watch()

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
