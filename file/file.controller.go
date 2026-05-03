package file

import (
	"forum/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

type getFileHandler struct {
	service *Service
}

func newGetFileHandler() http.Handler {
	return &getFileHandler{
		service: NewService(),
	}
}

func (h *getFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error converting id to int:", err)
		utils.SendMessage(w, "Not found", 404)
		return
	}

	f, err := h.service.getById(id)
	if err != nil {
		log.Println("Error getting file: err")
		utils.SendMessage(w, "Error getting file", 500)
		return
	}

	io.Copy(w, f)
}
