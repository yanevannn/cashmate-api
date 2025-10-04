package routes

import (
	"cashmate-api/controllers"
	"cashmate-api/middlewares"

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
			r.Post("/activate", controllers.ActivateUserHandler)
			r.Post("/resend-activation", controllers.ResendActivateCodeHandler)
			r.Post("/forgot-password", controllers.ForgotPasswordHandler)
			r.Post("/reset-password", controllers.ResetPasswordHandler)
		})

		// User Routes
		r.Route("/user", func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware) // All user routes require authentication

			r.Get("/{id}", controllers.GetUserByIDHandler) // Any authenticated user can get user by ID

			// Only admin can delete users
			r.With(middlewares.RoleMiddleware("administrator")).Get("/", controllers.GetAllUsersHandler)
			r.With(middlewares.RoleMiddleware("administrator")).Delete("/{id}", controllers.DeleteUserHandler)
		})

		// Category Routes
		r.Route("/categories", func(r chi.Router) {
			// All category routes require authentication and either admin or member role
			r.Use(middlewares.AuthMiddleware)
			r.Use(middlewares.RoleMiddleware("administrator", "member"))

			r.Get("/", controllers.GetAllCategoriesHandler)
			r.Post("/", controllers.CreateCategoryHandler)
			r.Put("/{id}", controllers.UpdateCategoryHandler)
			r.Delete("/{id}", controllers.DeleteCategoryHandler)
		})

		//Transaction Routes
		r.Route("/transactions", func(r chi.Router) {
			// All transaction routes require authentication and either admin or member role
			r.Use(middlewares.AuthMiddleware)
			r.Use(middlewares.RoleMiddleware("administrator", "member"))

			r.Get("/", controllers.GetAllTransactionHandler)
			r.Post("/", controllers.CreateTransactionHandler)
			r.Get("/{id}", controllers.GetTransactionByIdHandler)
			r.Put("/{id}", controllers.UpdateTransactionHandler)
			r.Delete("/{id}", controllers.DeleteTransactionHandler)
		})
	})

}
