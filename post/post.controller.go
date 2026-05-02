package post

import (
	"forum/utils"
	"log"
	"net/http"
)

const (
	maxMemory = 32 << 20 // 32 MB
)

type createPostHandler struct {
	service *service
}

func newCreatePostHandler() http.Handler {
	return &createPostHandler{
		service: newService(),
	}
}

func (h *createPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxMemory)
	dto := createPostDto{
		title:    r.FormValue("title"),
		content:  r.FormValue("content"),
		files:    r.MultipartForm.File["files"],
		authorId: r.Context().Value("userId").(int),
	}
	err := h.service.create(dto.title, dto.content, dto.authorId)
	if err != nil {
		log.Println("Error creating post:", err)
		utils.SendMessage(w, "Failed to create new post", 500)
		return
	}
	utils.SendMessage(w, "Post created successfully!", 201)
}
