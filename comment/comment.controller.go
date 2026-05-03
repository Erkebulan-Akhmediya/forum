package comment

import (
	"forum/utils"
	"log"
	"net/http"
	"strconv"
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
	postIdStr := r.PathValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		log.Println("Error converting post id:", err)
		utils.SendMessage(w, "Invalid post id", 400)
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		log.Println("Error reading file:", err)
		utils.SendMessage(w, "Failed to read file", 400)
		return
	}

	dto := createPostCommentDto{
		content:  r.FormValue("content"),
		auhtorId: r.Context().Value("userId").(int),
		postId:   postId,
		file:     &f,
	}
	if err := h.service.createPost(&dto); err != nil {
		log.Println("Error creating post comment:", err)
		utils.SendMessage(w, "Failed to create post comment", 500)
		return
	}
	utils.SendMessage(w, "Successfully created comment", 201)
}
