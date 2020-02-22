package main

import "time"

//go:noinline
func fC() {
	time.Sleep(time.Millisecond)
}

//go:noinline
func fB() {
	fC()
}

//go:noinline
func fA() {
	fB()
}

func main() {

	for i := 0; i < 1000; i++ {
		fA()
	}
}
