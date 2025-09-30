package routes

import (
	"cashmate-api/controllers"

	"github.com/go-chi/chi/v5"
)

const apiV1 = "/v1"

func RegisterRoutes(r *chi.Mux) {

	r.Route(apiV1, func(r chi.Router) {
		// Auth Routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", controllers.RegisterHandler)
			r.Post("/login", controllers.LoginHandler)
			r.Post("/refresh", controllers.RefreshTokenHandler)
		})
		
		// User Routes
		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}", controllers.GetUserByIDHandler)
			r.Get("/", controllers.GetAllUsersHandler)
			r.Delete("/{id}", controllers.DeleteUserHandler)
		})

		// Category Routes
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", controllers.GetAllCategoriesHandler)
			r.Post("/", controllers.CreateCategoryHandler)
			r.Put("/{id}", controllers.UpdateCategoryHandler)
			r.Delete("/{id}", controllers.DeleteCategoryHandler)
		})

		//Transaction Routes
		r.Route("/transactions", func(r chi.Router) {
			r.Get("/", controllers.GetALlTransactionHandler)
			r.Post("/", controllers.CreateTransactionHandler)
			r.Get("/{id}", controllers.GetTransactionByIdHandler)
			r.Put("/{id}", controllers.UpdateTransactionHandler)
			r.Delete("/{id}", controllers.DeleteTransactionHandler)
		})
	})

}
