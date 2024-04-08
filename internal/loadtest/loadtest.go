package loadtest

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type LoadTest struct {
	url        *url.URL
	req        int
	concurrent int
}

func NewTest(url *url.URL, request int, concurrent int) *LoadTest {
	return &LoadTest{
		url:        url,
		req:        request,
		concurrent: concurrent,
	}
}

func (lt *LoadTest) Run() {
	startTime := time.Now()
	ctx := context.Background()
	defer ctx.Done()

	wg := &sync.WaitGroup{}

	reqByWorker := lt.req / lt.concurrent

	for i := range lt.concurrent {
		wg.Add(1)
		println("add worker: ", i)
		go func() {
			defer wg.Done()
			worker(i, lt.url, reqByWorker)
		}()
	}

	wg.Wait()

	log.Println(time.Since(startTime))
}

func worker(id int, url *url.URL, totalReq int) {
	for i := range totalReq {
		log.Println("worker:", id, "request:", i)
		log.Println("response code:", makeRequest(url))
	}
}

func makeRequest(url *url.URL) int {

	response, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
		return response.StatusCode
	}
	return response.StatusCode
}
