package chcontext

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCdoe int) {
	s.written = true
}

//func (s *SpyStore) assertWasCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Error("store was not told to cancel")
//	}
//}
//
//func (s *SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Error("store was told to cancel")
//	}
//}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		// this simulates a slow fetch process by just iterating over the string
		// and adding in 10 ms intervalls to the res string
		// if ctx.Done() is signalled the bottom select will return error from the Fetch() function
		// And the goRoutine will also stop to prevent resource leak and do unnecessary work
		// if it finishes it will return the string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Print("spy store got cancelled")
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

//func (s *SpyStore) Cancel() {
//	s.cancelled = true
//}

func TestServer(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		//		store.assertWasNotCancelled()
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// cancel is a function
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// call cancel function after 5 milliseconds
		time.AfterFunc(5*time.Millisecond, cancel)
		//this changes the context of the request
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}

		//		if !store.cancelled {
		//			t.Errorf("store was not told to cancel")
		//		}
		//
		//		store.assertWasCancelled()
	})

}
