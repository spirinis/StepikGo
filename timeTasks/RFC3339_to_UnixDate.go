package timeTasks

import (
	"bufio"
	"os"
	"strings"
	"time"
)

func RFC3339_to_UnixDate() {
	// put your code here
	buf, err := bufio.NewReader(os.Stdin).ReadString('\n') // 1986-04-16T05:20:00+06:00
	if err != nil {
		panic(err)
	}
	buf = strings.TrimSpace(buf)
	t, err := time.Parse(time.RFC3339, buf)
	if err != nil {
		panic(err)
	}
	out := t.Format(time.UnixDate)
	_, err = os.Stdout.Write([]byte(out))
	if err != nil {
		panic(err)
	}
	//Wed Apr 16 05:20:00 +0600 1986

}
