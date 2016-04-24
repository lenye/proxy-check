package check

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	_ "github.com/Bogdan-D/go-socks4" // adds socks4 support
)

// Check makes channel with checking results
func Check(in chan *url.URL, out chan *Result, maxDelay time.Duration, target *url.URL, checkText string, workerCount int) {
	var done int64
	for w := 1; w <= workerCount; w++ {
		go func() {
			for n := range in {
				out <- query(n, maxDelay, target, checkText)
			}
			atomic.AddInt64(&done, 1)
			if done == int64(workerCount) {
				close(out)
			}
		}()
	}
}

func query(url *url.URL, maxDelay time.Duration, target *url.URL, checkText string) *Result {
	start := time.Now()
	res := &Result{
		URL: url}
	client := http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(url)},
		Timeout:   maxDelay}
	resp, err := client.Get(target.String())
	if err != nil {
		res.Err = err
		return res
	}
	diff := time.Now().UnixNano() - start.UnixNano()
	res.Delay = time.Duration(diff)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		return res
	}
	if !strings.Contains(string(body), checkText) {
		res.Err = errors.New("Expected text not found")
		return res
	}
	return res
}
