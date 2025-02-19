package timeTasks

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func UnixTime() {
	const now = 1589570165
	buf, err := bufio.NewReader(os.Stdin).ReadString('\n') // 12 мин. 13 сек.
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	buf = strings.TrimSpace(buf)
	parts := strings.Split(buf, " ")
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	sec, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}
	unTime := time.Unix(int64(now+60*min+sec), 0)
	unTime = unTime.UTC()
	out := unTime.Format(time.UnixDate)
	_, err = os.Stdout.Write([]byte(out)) // Fri May 15 19:28:18 UTC 2020
	if err != nil {
		panic(err)
	}

}
