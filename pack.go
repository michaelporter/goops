package main

import (
	"fmt"
	"runtime"
	"time"
)

var numberOfCrates int = 1000 // adjust this number to see perf at scale
var completed int

var readyForPrep chan *GrainBag
var readyForPrimary chan *GrainBag
var readyForSecondary chan *GrainBag
var readyForPlastic chan *GrainBag

func main() {
	runtime.GOMAXPROCS(3)
	jeffs := []*Jeff{new(Jeff), new(Jeff), new(Jeff)}

	readyForPrep = make(chan *GrainBag, 75)
	readyForPrimary = make(chan *GrainBag, 75)
	readyForSecondary = make(chan *GrainBag, 75)
	readyForPlastic = make(chan *GrainBag, 75)

	go prepareGrainBag(jeffs[0])
	go packPrimary(jeffs[0])
	go packSecondary(jeffs[1])
	go packInPlastic(jeffs[2])

	for c := 0; c < numberOfCrates; c++ {
		gb := new(GrainBag)
		readyForPrep <- gb
	}

	time.Sleep(10 * time.Second) // shim, should be checking for a full complete

	for j := 0; j < len(jeffs); j++ {
		fmt.Println("Jeff Ops ", j, jeffs[j].Ops)
	}
}
