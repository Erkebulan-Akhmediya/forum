package comment

import (
	"encoding/json"
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

	_, fh, err := r.FormFile("file")
	if err != nil {
		log.Println("Error reading file:", err)
		utils.SendMessage(w, "Failed to read file", 400)
		return
	}

	dto := createPostCommentDto{
		content:  r.FormValue("content"),
		auhtorId: r.Context().Value("userId").(int),
		postId:   postId,
		file:     fh,
	}
	if err := h.service.createPostComment(&dto); err != nil {
		log.Println("Error creating post comment:", err)
		utils.SendMessage(w, "Failed to create post comment", 500)
		return
	}
	utils.SendMessage(w, "Successfully created comment", 201)
}

type getAllByPostIdHandler struct {
	service *service
}

func newGetAllByPostIdHandler() http.Handler {
	return &getAllByPostIdHandler{
		service: newService(),
	}
}

func (h *getAllByPostIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page := utils.GetPage(r)
	postIdStr := r.PathValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		log.Println("Error parsing post id:", err)
		utils.SendMessage(w, "Invalid post id", 400)
		return
	}

	comments, err := h.service.getAllByPostId(postId, page)
	if err != nil {
		log.Println("Error getting post comments:", err)
		utils.SendMessage(w, "Failed to get post comments", 500)
		return
	}

	var dtos []*getPostCommentDto
	for _, c := range comments {
		authorDto := authorDto{
			Id:       c.author.id,
			Username: c.author.username,
			Email:    c.author.email,
		}
		commentDto := getPostCommentDto{
			Id:      c.id,
			Content: c.content,
			Author:  authorDto,
		}
		if c.fileId.Valid {
			commentDto.FileId = int(c.fileId.Int64)
		}
		dtos = append(dtos, &commentDto)
	}
	if dtos == nil {
		dtos = []*getPostCommentDto{}
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		log.Println("Error sendign comments:", err)
		utils.SendMessage(w, "Failed to send comments", 500)
		return
	}
}
