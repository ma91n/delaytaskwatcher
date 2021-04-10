package delaytaskwatcher

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var logger = log.Printf

const thresholdTerm = 3 * time.Second

type Options struct {
	Limit time.Duration
}

func Watch(opt ...Options) {
	term := thresholdTerm
	if len(opt) > 0 {
		term = opt[0].Limit
	}

	time.AfterFunc(term, func() {
		logger("task process time error: %s has passed", term)
		_ = pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	})
}
