package main

import (
	"testing"
	"time"
)

func TestSum(t testing.T) {
	chA := make(chan int)
	chB := make(chan int)

	res := sum(chA, chB)

	chA <- 1
	chB <- 2

	select {
	case su := <- res:
		if su != 3 {
			t.Errorf("incorrect sum, expected %d, received %d", 3, su)
		}
	case <-time.After(1000):
		t.Error("did not receive sum")
	}
}
