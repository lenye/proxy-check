package check

import (
	"net/url"
	"time"
)

// Result of proxy checking
type Result struct {
	Err   error
	Delay time.Duration
	URL   *url.URL
}
