package routes

import (
	"cashmate-api/controllers"

	"github.com/go-chi/chi/v5"
)

const apiV1 = "/v1"

func RegisterRoutes(r *chi.Mux) {

	r.Route(apiV1, func(r chi.Router) {
		// User Routes
		r.Route("/user", func(r chi.Router) {
			r.Post("/", controllers.CreateUserHandler)
			r.Get("/{id}", controllers.GetUserByIDHandler)
			r.Get("/", controllers.GetAllUsersHandler)
			r.Delete("/{id}", controllers.DeleteUserHandler)
		})
	})

}
