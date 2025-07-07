package HttpServer

import (
	"REST-API-pet-proj/Internal/HttpServer/Api/User"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"github.com/go-chi/chi/v5"
)

func InitRouter(storage *Sqlite.Storage) *chi.Mux {
	router := chi.NewRouter()
	//router.Use(middleware.RequestID)
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.URLFormat)
	//router.Use(middleware.Logger)
	//router.Use(middleware.RealIP)

	//позже добавлю мидлевары

	router.Route("/api", func(r chi.Router) {
		r.Post("/register", User.UserRegistrationHandler(storage))
		r.Post("/login", User.UserLoginHandler(storage))
		//r.Post("/like", )
		//r.Post("/post", )

		r.Get("/{username}", User.GetUserData(storage))
	})
	return router
}
