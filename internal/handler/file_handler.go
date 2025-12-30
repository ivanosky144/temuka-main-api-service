package handler

import (
	"net/http"

	"github.com/temuka-api-service/internal/service"
	rest "github.com/temuka-api-service/util/rest"
)

type FileHandler struct {
	fileService service.FileService
}

func NewFileHandler(service service.FileService) *FileHandler {
	return &FileHandler{fileService: service}
}

func (h *FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Could not parse multipart form"})
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Could not get uploaded file"})
		return
	}
	defer file.Close()

	// delegate to service
	url, err := h.fileService.UploadFile(r.Context(), header.Filename, file)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response := map[string]any{
		"message": "File has been uploaded",
		"url":     url,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}
