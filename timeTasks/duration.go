package timeTasks

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

func parse(str string) (t time.Time) {
	t, err := time.Parse("02.01.2006 15:04:05", str)
	if err != nil {
		panic(err)
	}
	return
}

func Duration() {
	buf, err := bufio.NewReader(os.Stdin).ReadString('\n') // 13.03.2018 14:00:15,12.03.2018 14:00:15
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	buf = strings.TrimSpace(buf)
	times := strings.Split(buf, ",")
	past := parse(times[0])
	future := parse(times[1])
	if past.After(future) {
		past, future = future, past
	}
	dur := future.Sub(past)
	_, err = os.Stdout.Write([]byte(dur.String())) // 24h0m0s
	if err != nil {
		panic(err)
	}

}
