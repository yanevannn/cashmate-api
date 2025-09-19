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

		// Category Routes
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", controllers.GetAllCategoriesHandler)
			r.Post("/", controllers.CreateCategoryHandler)
			r.Put("/{id}", controllers.UpdateCategoryHandler)
			// r.Delete("/{id}", controllers.DeleteCategoryHandler) // Soon: Implement if needed
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
