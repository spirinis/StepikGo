package find_file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var file_count, dir_count, max_size, prev_size, size3, size4 int

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		dir_count++
	}
	file_count++
	size := int(info.Size())

	if strings.Contains(info.Name(), ".txt") {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		r := csv.NewReader(file)
		data, err := r.ReadAll()

		if err != nil {
			fmt.Println(err)
			return errors.New("ошибка в walkfunc")
		}
		if len(data[0]) > 1 {
			fmt.Println(data)
			fmt.Println(path)
			fmt.Println("Файл найден. Размер ", size)
		}
	}

	if size >= max_size && size > prev_size {
		//prev_size = max_size
		max_size = size
	} else if size >= prev_size {
		prev_size = size
	} else if size >= size3 {
		size3 = size
	} else if size >= size4 {
		size4 = size
	}
	return nil
}

func File_finder_task() {
	const root = "read_file_tasks/find_file/files"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
	fmt.Println(dir_count, file_count)
	fmt.Println(max_size, prev_size, size3, size4)
	//io.WriteString(os.Stdout, strconv.Itoa(sum))
}
