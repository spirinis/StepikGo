package webservers

import (
	"fmt"
	"net/http"
)

// Напишите веб-сервер который по пути /api/user приветствует пользователя:
// Принимает и парсит параметр name и делает ответ "Hello,<name>!" порт :9000
func ServerQuery() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Write([]byte(fmt.Sprintf("Hello,%s!", name)))
}
