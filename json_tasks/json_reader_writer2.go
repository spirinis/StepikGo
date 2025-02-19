package json_tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jsonStruct2 []struct {
	Id int `json:"global_id"`
}

func Json_reader_writer2() {
	file, fo_err := os.Open("json_tasks/data-20190514T0100.json")
	if fo_err != nil {
		fmt.Println("Ошибка открытия ", fo_err)
		panic(fo_err)
	}
	defer file.Close()
	data, rf_err := io.ReadAll(file) //os.Stdin
	if rf_err != nil {
		fmt.Println("Ошибка чтения ", rf_err)
		panic(rf_err)
	}

	var myJsonStruct2 jsonStruct2
	if pj_err := json.Unmarshal(data, &myJsonStruct2); pj_err != nil {
		fmt.Println("Ошибка разбора ", pj_err)
		panic(pj_err)
	}
	sum := 0
	for _, id := range myJsonStruct2 {
		sum += id.Id
	}

	fmt.Println(sum)
}
