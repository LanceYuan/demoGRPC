package engine

import (
	"crawler/fetch"
	"log"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := worker(r)
		if err != nil {
			log.Printf("Worker err: %s\n", err)
			continue
		}
		requests = append(requests, parserResult.Request...)
		for _, item := range parserResult.Items {
			log.Printf("Get item: %s\n", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching url: %s\n", r.Url)
	body, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch: %s err: %v\n", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
