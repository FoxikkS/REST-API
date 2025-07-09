package UserActive

import (
	"REST-API-pet-proj/Internal/HttpServer/Api/User/Handlers"
	"REST-API-pet-proj/Internal/Storage"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"REST-API-pet-proj/Models"
	"net/http"
)

func UserPost(storage *Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Models.Post

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
