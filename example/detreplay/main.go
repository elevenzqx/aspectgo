package main

import (
	"github.com/elevenzqx/aspectgo/example/detreplay/worker"
)

func main() {
	n := 16
	chans := make([]chan struct{}, n)
	for i := 0; i < n; i++ {
		w := worker.W{X: i}
		chans[i] = w.DoWork()
	}
	for i := 0; i < n; i++ {
		<-chans[i]
	}
}
