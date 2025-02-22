package network

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).
func RequestHTTP() {
	// http запрос с методом GET
	resp, err := http.Get("http://127.0.0.1:8080/get") // https://golang.org https://go.dev/
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close() // закрываем тело ответа после работы с ним
	fmt.Println("Отправлен GET запрос")

	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", data) // печатаем ответ как строку
	//Ответ от сервера печатать как строку
}
