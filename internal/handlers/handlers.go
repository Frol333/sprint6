package handlers

import (
	"io"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final-tpl/internal/service"
)

// IndexHandler обрабатывает запросы к корневому эндпоинту и возвращает HTML из файла index.html.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла и конвертацию текста/морзе.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusBadRequest)
		return
	}

	converted, err := service.DetectAndConvert(string(fileBytes))
	if err != nil {
		http.Error(w, "failed to convert", http.StatusUnprocessableEntity)
		return
	}

	w.Write([]byte(converted))
}
