# delaytaskwatcher
goroutine delay task checker

## Usage

Installation.

```bash
go get github.com/ma91n/delaytaskwatcher
```

Use delaytaskwatcher like below with go func.

```go
package main

import (
	"github.com/ma91n/delaytaskwatcher"
)

func main() {
    go delaytaskwatcher.Watch()
    
    heavyTask()
}
```

Output is below.

```
>go run exmaple.go
2021/04/10 13:21:38 task process time error: 3s has passed
goroutine profile: total 2
1 @ 0xd2b01a 0xd56405 0xdd24f9 0xdd24e6 0xdd2432 0xd2abf6 0xd592a1
#       0xd56404        time.Sleep+0xe4         runtime/time.go:193
#       0xdd24f8        main.heavyTaskMain+0x98 github.com/ma91n/delaytaskwatcher/example/exmaple.go:24
#       0xdd24e5        main.heavyTask+0x85     github.com/ma91n/delaytaskwatcher/example/exmaple.go:19
#       0xdd2431        main.main+0x51          github.com/ma91n/delaytaskwatcher/example/exmaple.go:13
#       0xd2abf5        runtime.main+0x255      runtime/proc.go:225

1 @ 0xd53725 0xdc6ab5 0xdc6867 0xdc3338 0xdd23a5 0xd592a1
#       0xd53724        runtime/pprof.runtime_goroutineProfileWithLabels+0x64   runtime/mprof.go:716
#       0xdc6ab4        runtime/pprof.writeRuntimeProfile+0xd4                  runtime/pprof/pprof.go:724
#       0xdc6866        runtime/pprof.writeGoroutine+0xa6                       runtime/pprof/pprof.go:684
#       0xdc3337        runtime/pprof.(*Profile).WriteTo+0x3f7                  runtime/pprof/pprof.go:331
#       0xdd23a4        github.com/ma91n/delaytaskwatcher.Watch.func1+0x104     github.com/ma91n/delaytaskwatcher/watcher.go:26
```

You can know bottleneck where in your code.

## Options

```go
    go delaytaskwatcher.Watch(delaytaskwatcher.Options{
        Limit: 10 *time.Second,
        })
```

