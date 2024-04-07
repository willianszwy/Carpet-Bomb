package loadtest

import (
	"context"
	"log"
	"net/http"
	"net/url"
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

	for i := range lt.req {
		log.Println(i)
		doRequest(lt.url)
	}

	log.Println(time.Since(startTime))
}

func doRequest(url *url.URL) {

	response, err := http.Get(url.String())
	if err != nil {
		log.Println(err)
		return
	}
	println(response.StatusCode)
}
