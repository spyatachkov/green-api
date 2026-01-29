package recovery

import (
	"log"
	"runtime/debug"
)

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("goroutine panic recovered: %v\n%s", err, debug.Stack())
			}
		}()
		fn()
	}()
}

func Recover() {
	if err := recover(); err != nil {
		log.Printf("panic recovered: %v\n%s", err, debug.Stack())
	}
}
