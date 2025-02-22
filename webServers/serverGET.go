package webservers

import (
	"fmt"
	"net/http"
)

// Напишите веб сервер, который по пути /get отдает текст "Hello, web!".
// Порт должен быть :8080.
func ServerGET() {
	// Регистрируем обработчики для разных путей
	http.HandleFunc("/get", handleRequest)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("Запуск сервера")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// В зависимости от метода HTTP-запроса вызываем соответствующий обработчик
	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// Обработчик для GET-запросов
func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен GET запрос", r)
	fmt.Fprintln(w, "Это GET-запрос!")
	w.Write([]byte("Hello, web!"))
}
