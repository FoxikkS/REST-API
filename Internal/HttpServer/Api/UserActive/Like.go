package UserActive

import (
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"net/http"
)

func PutALike(storage *Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
