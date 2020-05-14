package responses

import "net/http"

type OkayResponse struct {
	Status string `json:"status"`
}

func (response OkayResponse) SendResponse(w http.ResponseWriter, r *http.Request) {
	response.Status = OK
	w.WriteHeader(http.StatusOK)
	sendResponse(w, r, response)
}
