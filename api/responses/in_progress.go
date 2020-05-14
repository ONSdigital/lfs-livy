package responses

import (
	"net/http"
	"time"
)

type InProgressResponse struct {
	Status  string `json:"status"`
	When    string `json:"time"`
	Message string `json:"message"`
}

func (response InProgressResponse) SendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = OK
	response.When = time.Now().String()
	response.Message = "job submitted"
	w.WriteHeader(http.StatusAccepted)
	sendResponse(w, r, response)
}
