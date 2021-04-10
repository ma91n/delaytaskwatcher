package delaytaskwatcher

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var logger = log.Printf

const thresholdTerm = 3 * time.Second

func Watch() {
	time.AfterFunc(thresholdTerm, func() {
		logger("task process time error: %s has passed", thresholdTerm)
		_ = pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	})
}
