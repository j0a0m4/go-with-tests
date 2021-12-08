package store

import (
	"context"
	"errors"
	assert "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Write([]byte) (n int, err error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) WriteHeader(int) {
	s.written = true
}

type SpyStore struct {
	response string
	written  bool
	t        *testing.T
}

func (s *SpyStore) Cancel() {
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
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

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		is := assert.New(t)
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)
		is.Equal(data, res.Body.String())
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		is := assert.New(t)
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := &SpyResponseWriter{}

		svr.ServeHTTP(res, req)
		is.True(!res.written)
	})
}
