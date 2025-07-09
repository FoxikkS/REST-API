package UserActive

import (
	"REST-API-pet-proj/Internal/HttpServer/Api/User/Handlers"
	"REST-API-pet-proj/Internal/Storage"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"net/http"
)

func UserPost(storage *Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username" validate:"required"`
			Title    string `json:"title" validate:"required,min=3,max=100"`
			Content  string `json:"content" validate:"required,min=10"`
		}

		if !Handlers.ParseAndValidateJSON(w, r, &req) {
			return
		}

		err := Storage.CreatePost(req.Username, req.Title, req.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"post created"}`))
	}
}
