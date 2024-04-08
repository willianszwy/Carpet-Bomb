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
	url            *url.URL
	req            int
	concurrent     int
	ResponseStatus map[int]int
	mutex          sync.Mutex
}

func NewTest(url *url.URL, request int, concurrent int) *LoadTest {
	if concurrent < 1 {
		concurrent = 1
	}

	if request < 1 {
		request = 1
	}

	if concurrent > request {
		concurrent = request
	}

	return &LoadTest{
		url:            url,
		req:            request,
		concurrent:     concurrent,
		ResponseStatus: map[int]int{},
	}
}

func (lt *LoadTest) Run() {
	startTime := time.Now()
	ctx := context.Background()
	defer ctx.Done()

	wg := &sync.WaitGroup{}

	reqByWorker := lt.req / lt.concurrent
	remainder := lt.req % lt.concurrent

	for i := range lt.concurrent {
		wg.Add(1)
		println("add worker: ", i)
		go func() {
			defer wg.Done()
			tr := reqByWorker
			if i == (lt.concurrent - 1) {
				tr += remainder
			}
			worker(i, lt.url, tr, lt.salveResponses)
		}()
	}

	wg.Wait()

	println("Results:")
	println("Total time:", time.Since(startTime).String())
	println("Total requests:", lt.req)

	total200 := lt.ResponseStatus[200]
	println("Total 200 request:", total200)

	for key, value := range lt.ResponseStatus {
		println("status:", key, "total:", value)
	}

}

func worker(id int, url *url.URL, totalReq int, counter func(int)) {
	log.Println("worker:", id, "total request:", totalReq)
	for i := range totalReq {
		log.Println("worker:", id, "request:", i)
		statusCode := makeRequest(url)
		log.Println("response code:", statusCode)
		counter(statusCode)
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

func (lt *LoadTest) salveResponses(statusCode int) {
	lt.mutex.Lock()
	defer lt.mutex.Unlock()

	count, exists := lt.ResponseStatus[statusCode]

	if !exists {
		count = 0
	}
	lt.ResponseStatus[statusCode] = count + 1
}
