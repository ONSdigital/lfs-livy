package responses

import "net/http"

type UnknownJob struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func (response UnknownJob) SendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = Error
	w.WriteHeader(http.StatusBadRequest)
	sendResponse(w, r, response)
}
