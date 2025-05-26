// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	counter int
	lock    = sync.Mutex{}
)

func main() {
	ticker := time.NewTicker(time.Duration(30+rand.Int31n(30)) * time.Second)
	ticker2 := time.NewTicker(time.Duration(30+rand.Int31n(30)) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				lock.Lock()
				counter++
				lock.Unlock()

				log.Println("fetch the", counter, "times")

				if fetch() {
					log.Println("Found!")
					notify()
					continue
				}

				log.Println("not found, waiting for the next fetch...")
			}
		}
	}()

	for {
		select {
		case <-ticker2.C:
			lock.Lock()
			counter++
			lock.Unlock()

			log.Println("fetch the", counter, "times")

			if fetch() {
				log.Println("Found!")
				notify()
				continue
			}

			log.Println("not found, waiting for the next fetch...")
		}
	}

}
