package comment

import (
	"forum/auth"
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	commentHandler := utils.MethodHandler{
		http.MethodPost: auth.NewMiddleware(newCreateHandler()),
	}
	http.Handle("/post/{postId}/comment", commentHandler)
}
