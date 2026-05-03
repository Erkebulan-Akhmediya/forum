package post

import (
	"forum/utils"
	"log"
	"net/http"
)

const (
	maxMemory = 32 << 20 // 32 MB
)

type createHandler struct {
	service *service
}

func newCreateHandler() http.Handler {
	return &createHandler{
		service: newService(),
	}
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxMemory)
	dto := createDto{
		title:    r.FormValue("title"),
		content:  r.FormValue("content"),
		files:    r.MultipartForm.File["files"],
		authorId: r.Context().Value("userId").(int),
	}
	if err := h.service.create(&dto); err != nil {
		log.Println("Error creating post:", err)
		utils.SendMessage(w, "Failed to create new post", 500)
		return
	}
	utils.SendMessage(w, "Post created successfully!", 201)
}
