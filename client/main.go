package main 

import (
	"fmt"
	"flag"
	"sync"
	"time"
)

//	Keyspace 	= demo
//	Table		= cas(key int, id text, ver int, val text)

var (
	ID string
	mutex = &sync.Mutex{}
	servers = []string{"128.52.162.127:5001", "128.52.162.122:500`", "128.52.162.123:5001"}	
)

// mark the phase
const QUERY=0
const PREWRITE=1
const RFINALIZE=2
const WFINALIZE=3

func main() {
	var num int
	flag.IntVar(&num, "num of RW", 1000, "numofRW")
	flag.StringVar(&ID, "clientID", "172.17.0.1", "input client ID")
	flag.Parse()		

	test(num)
	// client()
}

func test(num int) {
	wTime := make(chan time.Duration)
	rTime := make(chan time.Duration)
	var WTotal, RTotal int = 0,0

	for i := 0; i < num; i++ {
		go testW(wTime)
		go testR(rTime)
	}

	for i := 0; i < num; i++ {
		WTotal += int(<-wTime/time.Millisecond)
		RTotal += int(<-rTime/time.Millisecond)
	}

	fmt.Printf("Avg write time: %f ms\n", float64(WTotal)/float64(num))
	fmt.Printf("Avg read time: %f ms\n", float64(RTotal)/float64(num))
}

func testW(wTime chan time.Duration){
	bytes := make([]byte, 2048, 2048)
	mutex.Lock()
	start := time.Now()
	write("0", bytes)
	end := time.Now()
	mutex.Unlock()
	wTime <- end.Sub(start)
}

func testR(rTime chan time.Duration){
	mutex.Lock()
	start := time.Now()
	read("0")
	end := time.Now()
	mutex.Unlock()
	rTime <- end.Sub(start)
}