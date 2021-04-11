package delaytaskwatcher

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

var logger = log.Printf

const thresholdTerm = 3 * time.Second

type Options struct {
	// Stack trace output limit term
	Limit time.Duration

	// - debug=0 writes the gzip-compressed protocol buffer described.
	// - debug=1 writes the legacy text format with comments
	// translating addresses to function names and line numbers.
	// - debug=2 means to print the goroutine stacks in the same form that a Go program uses
	// when dying due to an unrecovered panic.
	// default is 1.
	Debug *int

	// NOW DEVELOP OPTION!
	// 出力対象のGoroutineID
	TargetGoID *int
}

func NewCurrentGoroutine() Options {
	return Options{
		TargetGoID: intp(goid()),
	}
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func intp(i int) *int {
	return &i
}

func Watch(opts ...Options) (stop func()) {
	term := thresholdTerm
	debug := 1
	var goID *int

	for _, opt := range opts {
		if opt.Limit > 0 {
			term = opt.Limit
		}

		if debugOpt := opt.Debug; debugOpt != nil {
			debug = *debugOpt
		}

		if goIDOpt := opt.TargetGoID; goIDOpt != nil {
			goID = goIDOpt
		}

	}

	afterFunc := time.AfterFunc(term, func() {
		fmt.Println("goroutine ID", goid())

		if goID != nil {
			//lookup
		}

		var buf bytes.Buffer
		_ = pprof.Lookup("goroutine").WriteTo(&buf, debug)

		logger("task process time error: %s has passed: %s", thresholdTerm, buf.String())
	})

	return func() {
		afterFunc.Stop()
	}

}
