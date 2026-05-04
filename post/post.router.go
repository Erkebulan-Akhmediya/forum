package post

import (
	"forum/auth"
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	postHandler := utils.MethodHandler{
		http.MethodGet:  newGetAllHandler(),
		http.MethodPost: auth.NewMiddleware(newCreateHandler()),
	}
	http.Handle("/post", postHandler)

	onePostHandler := utils.MethodHandler{
		http.MethodPost: auth.NewMiddleware(newReactHandler()),
	}
	http.Handle("/post/{id}/react", onePostHandler)
}
