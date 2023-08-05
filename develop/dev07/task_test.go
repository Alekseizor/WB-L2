package main

import (
	"log"
	"testing"
	"time"
)

func TestComposition(t *testing.T) {
	sig := func(after time.Duration, msg string) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			c <- msg
			time.Sleep(after)
		}()
		return c
	}

	result := composition(
		sig(1*time.Second, "hello"),
		sig(1*time.Second, "wb"),
	)
	for res := range result {
		log.Println(res)
	}
}
