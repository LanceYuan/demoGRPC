package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigWorkChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) ParserResult {
	in := make(chan Request)
	out := make(chan ParserResult)
	c.Scheduler.ConfigWorkChan(in)
	for i := 0; i <= c.WorkCount; i++ {
		createWorker(in, out)
	}
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get item: %s\n", item)
		}
		for _, request := range result.Request {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
