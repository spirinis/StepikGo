package json_tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jsonStruct struct {
	Students []struct {
		Rating []int
	}
}

func Json_reader_writer() {
	file, fo_err := os.Open("json_tasks/data.json")
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

	var myJsonStruct jsonStruct
	if pj_err := json.Unmarshal(data, &myJsonStruct); pj_err != nil {
		fmt.Println("Ошибка разбора ", pj_err)
		panic(pj_err)
	}
	sum := 0
	for _, stud := range myJsonStruct.Students {
		sum += len(stud.Rating)
	}
	out := float64(sum) / float64(len(myJsonStruct.Students))

	type ansStruct struct {
		Average float64
	}
	ans := ansStruct{out}
	data, mj_err := json.MarshalIndent(ans, "", "    ")
	if mj_err != nil {
		fmt.Println("Ошибка создания ", mj_err)
		panic(mj_err)
	}
	_, err := os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
