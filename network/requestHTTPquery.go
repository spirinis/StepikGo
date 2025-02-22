package network

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func RequestHTTPquery() {
	// Считывание параметров
	var name, age string
	fmt.Scan(&name, &age)
	// Создаем URL с параметрами
	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	fullURL := baseURL + "?" + params.Encode()

	// http запрос с методом GET
	resp, err := http.Get(fullURL) // https://golang.org https://go.dev/
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close() // закрываем тело ответа после работы с ним

	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", data) // печатаем ответ как строку
}
