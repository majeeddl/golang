package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan string {
	ch := make(chan string)

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

// func Racer(a, b string) (winner string) {

// 	// startA := time.Now()
// 	// http.Get(a)
// 	// aDuration := time.Since(startA)

// 	// startB := time.Now()
// 	// http.Get(b)

// 	// bDuration := time.Since(startB)
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
