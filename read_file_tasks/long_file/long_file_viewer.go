package long_file

import (
	"bufio"
	"fmt"
	"os"
)

func Long_file_viewer() {
	const path = "read_file_tasks/long_file/long_file.data"
	file, fo_err := os.Open(path)
	if fo_err != nil {
		panic(fo_err)
	}
	defer file.Close()
	// reader := bufio.NewReader(file)
	// buf := make([]byte, 20)
	// i := 0
	// for _, r_err := reader.Read(buf); i < 10 && r_err == nil && r_err != io.EOF; _, r_err = reader.Read(buf) {
	// 	fmt.Println(buf)
	// 	i++
	// }
	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i, b := range data {
			if b == ';' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	n := 0
	for scanner.Scan() {
		n++
		token := scanner.Text()
		if token == "0" {
			fmt.Println(n)
		}
	}
}
