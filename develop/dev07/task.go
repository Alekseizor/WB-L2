package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func composition(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	// Создать горутины для чтения из каждого канала
	go func() {
		defer close(single)
		wg := &sync.WaitGroup{}
		for _, ch := range channels {
			wg.Add(1)
			go func(wg *sync.WaitGroup, c <-chan interface{}) {
				defer wg.Done()
				for v := range c {
					// Передать значение из канала
					single <- v
				}
			}(wg, ch)
		}
		wg.Wait()
	}()
	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	result := composition(
		sig(3*time.Second),
		sig(5*time.Second),
		sig(10*time.Second),
		sig(12*time.Second),
	)
	for res := range result {
		log.Println(res)
	}
	fmt.Printf("Done after %v\n", time.Since(start))
}
