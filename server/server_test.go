package server

import (
	"encoding/json"
	assert "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
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
	server := NewPlayerServer(&store)

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
	server := NewPlayerServer(&store)
	player := "Pepper"

	for i := 0; i < 3; i++ {
		req, _ := newPostWinRequest(player)
		server.ServeHTTP(httptest.NewRecorder(), req)
	}

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := newGetScoreRequest(player)

		server.ServeHTTP(res, req)

		is := assert.New(t)
		is.Equal(res.Code, http.StatusOK)
		is.Equal(res.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := newLeagueRequest()

		server.ServeHTTP(res, req)

		want := []Player{{"Pepper", 3}}
		var got []Player
		err := json.NewDecoder(res.Body).Decode(&got)

		is := assert.New(t)
		is.Equal(res.Code, http.StatusOK)
		is.Equal(res.Header().Get("content-type"), jsonContentType)
		is.Equal(got, want)
		is.NoErr(err)
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wanted := []Player{
			{"Cleo", 32},
			{"Christ", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{
			scores:   nil,
			winCalls: nil,
			league:   wanted,
		}
		server := NewPlayerServer(&store)

		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []Player

		err := json.NewDecoder(res.Body).Decode(&got)

		is := assert.New(t)
		is.NoErr(err)
		is.Equal(res.Code, http.StatusOK)
		is.Equal(got, wanted)
		is.Equal(res.Header().Get("content-type"), jsonContentType)
	})
}

func newPostWinRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, "/players/"+name, nil)
}

func newGetScoreRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/players/"+name, nil)
}

func newLeagueRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "/league", nil)
}
