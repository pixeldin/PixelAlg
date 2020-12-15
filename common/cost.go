package common

import (
	"log"
	"time"
)

func TimeCost(fname string) {
	s := time.Now()
	log.Printf("Func \"%v\" cost %v(ts).", fname, time.Since(s).Milliseconds())
}

func MockSome() {
	defer TimeCost("Mock")
	// do some...
	time.Sleep(1 * time.Second)
}
