package webservers

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
Напиши веб сервер (порт :3333) - счетчик который будет обрабатывать GET (/count) и POST (/count) запросы:
GET:  возвращает счетчик
POST: увеличивает ваш счетчик на значение  (с ключом "count") которое вы получаете из формы,
но если пришло НЕ число то нужно ответить клиенту: "это не число" со статусом http.StatusBadRequest (400).
*/
func ServerRequestCounter() {
	// Регистрируем обработчик для пути "/count"
	http.HandleFunc("/count", handlerCounter)
	// Запускаем веб-сервер на порту 3333
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

var counter int = 0

// Обработчик HTTP-запросов
func handlerCounter(w http.ResponseWriter, r *http.Request) {
	// проверяем что метод POST
	if r.Method == "POST" {
		r.ParseForm()
		str := r.Form.Get("count")
		count, err := strconv.Atoi(str)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		counter += count
	}
	if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(counter)))
	}
}
