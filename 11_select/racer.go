package chselect

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// <- waiting for a channel is a blocking call
// select allows to wait on multiple channels
// the first one to send a value "wins" and the code underneath the case is executed
// select is executed once if a channel is ready, if multiple ready choose random, if none ready block (non blocking if default case exist)
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		// time.After returns a chan time.Time
		// helps to make a timeout to unblock the select statement when no channel is returning anything
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//func Racer(a, b string) (winner string) {
//	aDuration := measureResponseTime(a)
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//	return b
//}
//
//func measureResponseTime(url string) time.Duration {
//	startA := time.Now()
//	http.Get(url)
//	return time.Since(startA)
//}
//
