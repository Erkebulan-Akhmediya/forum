package post

import (
	"encoding/json"
	"forum/utils"
	"log"
	"net/http"
)

const (
	maxMemory = 32 << 20 // 32 MB
)

type createHandler struct {
	service *Service
}

func newCreateHandler() http.Handler {
	return &createHandler{
		service: NewService(),
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

type getAllHandler struct {
	service *Service
}

func newGetAllHandler() http.Handler {
	return &getAllHandler{
		service: NewService(),
	}
}

func (h *getAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page := utils.GetPage(r)
	posts, err := h.service.getAll(page)
	if err != nil {
		log.Println("Error getting posts", err)
		utils.SendMessage(w, "Failed to get posts", 500)
		return
	}

	var dtos []*getDto
	for _, p := range posts {
		authorDto := authorDto{
			Id:       p.author.id,
			Username: p.author.username,
			Email:    p.author.email,
		}
		postDto := getDto{
			Id:      p.id,
			Author:  authorDto,
			Title:   p.title,
			Content: p.content,
			FileIds: p.fileIds,
		}
		dtos = append(dtos, &postDto)
	}

	if dtos == nil {
		dtos = make([]*getDto, 0)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		log.Println("Error encoding posts:", err)
		utils.SendMessage(w, "Failed to send posts", 500)
		return
	}
}
