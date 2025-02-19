package timeTasks

import (
	"bufio"
	"os"
	"strings"
	"time"
)

func ChangeDateTime() {
	// put your code here
	buf, err := bufio.NewReader(os.Stdin).ReadString('\n') // 2020-05-15 08:00:00
	if err != nil {
		panic(err)
	}
	buf = strings.TrimSpace(buf)
	t, err := time.Parse(time.DateTime, buf)
	if err != nil {
		panic(err)
	}
	if t.Hour() >= 13 {
		t = t.AddDate(0, 0, 1)
	}
	out := t.Format(time.DateTime)
	_, err = os.Stdout.Write([]byte(out))
	if err != nil {
		panic(err)
	}
	// 2020-05-15 08:00:00

}
