package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// responsJSON выполняет сериализацию объектов доменной области в JSON и отправляет ответ клиенту
func responsJSON(w http.ResponseWriter, data any, status int) {
	// Устанавливаем тип контента ответа
	w.Header().Set("Content-Type", "application/json")
	// Устанавливаем HTTP-статус ответа
	w.WriteHeader(status)

	// Кодируем данные в JSON и отправляем клиенту
	if err := json.NewEncoder(w).Encode(map[string]any{"result": data}); err != nil {
		log.Printf("[ERROR] responsJSON: %s\n", err.Error())
		// Логируем ошибку, если не удалось отправить ответ
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// responsErrorJSON выполняет сериализацию ошибки в JSON и отправляет ответ клиенту
func responsErrorJSON(w http.ResponseWriter, err error, status int) {
	// Устанавливаем тип контента ответа
	w.Header().Set("Content-Type", "application/json")
	// Устанавливаем HTTP-статус ответа
	w.WriteHeader(status)

	// Кодируем информацию об ошибке в JSON и отправляем клиенту
	if err := json.NewEncoder(w).Encode(map[string]any{"error": err.Error()}); err != nil {
		log.Printf("[ERROR] responsErrorJSON: %s\n", err.Error())
		// Логируем ошибку, если не удалось отправить ответ
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Дополнительно логируем ошибку как предупреждение
	log.Printf("[WARN] responsErrorJSON: %s\n", err.Error())
}
