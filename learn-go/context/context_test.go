package context1

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	// cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled(t *testing.T) {
	s.t.Helper()
	if !s.cancelled {
		t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled(t *testing.T) {
	s.t.Helper()
	if s.cancelled {
		t.Error("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {

		store := &SpyStore{response: data}

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelingCtx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancelingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		// if !store.cancelled {
		// 	t.Error("store was not told to cancel")

		// }

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasCancelled(t)
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCancelled(t)

	})

}
