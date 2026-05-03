package comment

import (
	"forum/auth"
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	commentHandler := utils.MethodHandler{
		http.MethodGet:  newGetAllByPostIdHandler(),
		http.MethodPost: auth.NewMiddleware(newCreateHandler()),
	}
	http.Handle("/post/{postId}/comment", commentHandler)
}
