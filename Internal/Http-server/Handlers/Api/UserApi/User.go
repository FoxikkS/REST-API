package UserApi

import (
	"REST-API-pet-proj/Internal/Storage"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type registration struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=4,max=31"`
	Email    string `json:"email" validate:"required,email"`
}

func UserRegistrationHandler(storage *Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var validate = validator.New()
		var req registration

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = validate.Struct(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = Storage.SaveUser(req.Username, req.Email, string(hash))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User Registration Successful"))
	}
}

func UserLoginHandler(storage *Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetUserHandler(storage Sqlite.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
