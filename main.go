package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Получение IP-адреса клиента
	ipAddress := r.RemoteAddr

	// Получение заголовков клиента
	headers := make(map[string]string)
	for key, values := range r.Header {
		// Если есть только одно значение, сохраняем его
		if len(values) == 1 {
			headers[key] = values[0]
		}
	}

	// Проверка наличия заголовка X-Forwarded-For
	if forwardedFor, ok := headers["X-Forwarded-For"]; ok {
		headers["IsProxy"] = fmt.Sprintf("true (X-Forwarded-For: %s)", forwardedFor)
	} else {
		headers["IsProxy"] = "false"
	}

	// Добавление IP-адреса в заголовки
	headers["IP"] = ipAddress

	// Кодирование заголовков в формат JSON
	jsonResponse, err := json.Marshal(headers)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}

	// Установка заголовка ответа
	w.Header().Set("Content-Type", "application/json")

	// Отправка JSON в ответ клиенту
	_, err = w.Write(jsonResponse)
	if err != nil {
	}
}

func main() {
	// Настройка обработчика HTTP-запросов
	http.HandleFunc("/", handler)

	// Запуск веб-сервера на порту 8080
	addr := ":8080"
	go func() {
		fmt.Printf("Сервер запущен на %s\n", addr)
	}()

	err := http.ListenAndServe(addr, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error listening for server: %s\n", err)
	}
}
