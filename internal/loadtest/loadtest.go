package loadtest

import "net/url"

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

}
