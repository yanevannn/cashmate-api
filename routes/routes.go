package routes

import (
	"cashmate-api/controllers"
	"net/http"
)

const apiV1 = "/v1"

func RegisterRoutes() {
	http.HandleFunc(apiV1+"/user", controllers.CreateUserHandler)
}
