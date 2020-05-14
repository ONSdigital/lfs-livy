package responses

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	Error   = "ERROR"
	OK      = "OK"
	Success = "SUCCESS"
)

type Response interface {
	SendResponse(w http.ResponseWriter, r *http.Request)
}

func sendResponse(w http.ResponseWriter, r *http.Request, response Response) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().
			Str("client", r.RemoteAddr).
			Str("uri", r.RequestURI).
			Msg("json.NewEncoder() failed")
	}
}
