package countdown

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer, sleeper Sleeper, start int, msg string) {
	for i := start; i > 0; i-- {
		displayOnTimeout(out, sleeper, i)
	}
	fmt.Fprint(out, msg)
}

func displayOnTimeout(out io.Writer, sleeper Sleeper, count int) {
	sleeper.Sleep()
	fmt.Fprintln(out, count)
}
