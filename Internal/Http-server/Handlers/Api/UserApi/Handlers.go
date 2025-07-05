package UserApi

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func DataValidator[T any](data *T) (*validator.Validate, error) {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		return nil, err
	}
	return v, err
}

func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ParseAndValidateJSON[T any](w http.ResponseWriter, r *http.Request, target *T) bool {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return false
		}
	}
	_, err = DataValidator(target)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}
