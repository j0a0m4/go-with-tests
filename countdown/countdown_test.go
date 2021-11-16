package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdown struct {
	Calls []string
	t     *testing.T
}

func (s *SpyCountdown) VerifyCalls(want []string) {
	s.t.Helper()
	if !reflect.DeepEqual(want, s.Calls) {
		s.t.Errorf("wanted calls %v but got %v", want, s.Calls)
	}
}

func (s *SpyCountdown) Sleep() {
	s.called(sleep)
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.called(write)
	return
}

func (s *SpyCountdown) called(operation string) {
	s.Calls = append(s.Calls, operation)
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountdown{t: t}
		Countdown(buffer, spy, 3, "GO!")

		got := buffer.String()
		want := `3
2
1
GO!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every count", func(t *testing.T) {
		spy := &SpyCountdown{t: t}
		Countdown(spy, spy, 3, "GO!")

		want := []string{
			sleep, write, sleep, write, sleep, write, write,
		}

		spy.VerifyCalls(want)
	})
}
