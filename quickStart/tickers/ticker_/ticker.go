package ratelimiting

import (
	"log"
	"time"
)

func Ticker() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Println("tick at ", t)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	log.Println("ticker stopped")
}
