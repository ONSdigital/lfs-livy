package responses

import "net/http"

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func (response ErrorResponse) SendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = Error
	w.WriteHeader(http.StatusBadRequest)
	sendResponse(w, r, response)
}
