package main

import (
	"log"
	"time"
)

func main() {
	// write a boilerplate code to crawl a website
	log.Println("start watching")
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("start finding")
			if fetch() {
				notify()
			}
		}
	}
}
