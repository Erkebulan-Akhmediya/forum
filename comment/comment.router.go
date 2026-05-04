package comment

import (
	"forum/auth"
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	postCommentHandler := utils.MethodHandler{
		http.MethodGet:  newGetAllByPostIdHandler(),
		http.MethodPost: auth.NewMiddleware(newCreateHandler()),
	}
	http.Handle("/post/{postId}/comment", postCommentHandler)

	replyCommentHandler := utils.MethodHandler{
		http.MethodPost: auth.NewMiddleware(newReplyHandler()),
	}
	http.Handle("/comment/{commentId}/comment", replyCommentHandler)
}
