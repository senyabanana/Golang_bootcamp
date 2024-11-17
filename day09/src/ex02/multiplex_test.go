package main

import (
	"testing"
)

func TestMultiplex(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		ch1 <- "from ch1"
		close(ch1)
	}()

	go func() {
		ch2 <- "from ch2"
		close(ch2)
	}()

	out := multiplex(ch1, ch2)

	received := map[string]bool{}

	for msg := range out {
		received[msg.(string)] = true
	}

	if !received["from ch1"] || !received["from ch2"] {
		t.Errorf("Expected messages from both channels, got %v", received)
	}
}

func TestMultiplexEmptyChannels(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	close(ch1)
	close(ch2)

	out := multiplex(ch1, ch2)

	_, ok := <-out
	if ok {
		t.Errorf("Expected closed channel")
	}
}
