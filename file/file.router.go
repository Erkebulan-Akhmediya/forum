package file

import (
	"forum/auth"
	"forum/utils"
	"net/http"
)

func RegisterRoutes() {
	fileHandler := utils.MethodHandler{
		http.MethodGet: auth.NewMiddleware(newGetFileHandler()),
	}
	http.Handle("/file/{id}", fileHandler)
}
