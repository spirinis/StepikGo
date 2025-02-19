package json_tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	Students []struct {
		Rating []int
	}
}

func Good_solution() {
	file, fo_err := os.Open("json_tasks/data.json")
	if fo_err != nil {
		fmt.Println("Ошибка открытия ", fo_err)
		panic(fo_err)
	}
	defer file.Close()
	// data, rf_err := io.ReadAll(file) //os.Stdin
	// if rf_err != nil {
	// 	fmt.Println("Ошибка чтения ", rf_err)
	// 	panic(rf_err)
	// }

	var group Group
	if err := json.NewDecoder(file).Decode(&group); err != nil { //os.Stdin
		panic(err)
	}

	var avg float64
	for _, st := range group.Students {
		avg += float64(len(st.Rating))
	}
	avg /= float64(len(group.Students))

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(struct{ Average float64 }{avg})
}
