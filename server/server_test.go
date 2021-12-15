package server

import (
	assert "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := PlayerServer{Store: &store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		req, _ := newGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		is := assert.New(t)
		is.Equal(res.Body.String(), "20")
		is.Equal(res.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req, _ := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		is := assert.New(t)
		is.Equal(res.Body.String(), "10")
		is.Equal(res.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req, _ := newGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		is := assert.New(t)
		is.Equal(res.Code, http.StatusNotFound)
	})

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		req, _ := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		is := assert.New(t)
		is.Equal(res.Code, http.StatusAccepted)
		is.Equal(len(store.winCalls), 1)
		is.Equal(store.winCalls[0], player)
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{map[string]int{}}
	server := PlayerServer{Store: &store}
	player := "Pepper"

	for i := 0; i < 3; i++ {
		req, _ := newPostWinRequest(player)
		server.ServeHTTP(httptest.NewRecorder(), req)
	}

	res := httptest.NewRecorder()
	req, _ := newGetScoreRequest(player)

	server.ServeHTTP(res, req)

	is := assert.New(t)
	is.Equal(res.Code, http.StatusOK)
	is.Equal(res.Body.String(), "3")
}

func newPostWinRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, "/players/"+name, nil)
}

func newGetScoreRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/players/"+name, nil)
}
