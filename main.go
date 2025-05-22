// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"log"
	"time"
)

var counter int

func main() {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			counter++
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
