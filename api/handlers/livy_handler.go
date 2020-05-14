package handlers

import (
	"encoding/json"
	"github.com/ONSDigital/lfs-livy/api/responses"
	"github.com/ONSDigital/lfs-livy/api/services"
	"github.com/rs/zerolog/log"
	"net/http"
)

type LivyJobHandler struct {
	services.LivyRequestService
}

type SubmitRequest struct {
	Job string
	Jar string
}

func NewLivyJobHandler() *LivyJobHandler {
	return &LivyJobHandler{services.LivyRequestService{}}
}

func (lj *LivyJobHandler) RunJobHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().
		Str("client", r.RemoteAddr).
		Str("uri", r.RequestURI).
		Msg("Received livy Job Request")

	var sr SubmitRequest

	err := json.NewDecoder(r.Body).Decode(&sr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if sr.Job == "" {
		log.Error().Msg("Run Job - job name not set")
		responses.ErrorResponse{ErrorMessage: "run job  - job name not set"}.SendResponse(w, r)
		return
	}

	if sr.Jar == "" {
		log.Error().Msg("Run Job - jar name not set")
		responses.ErrorResponse{ErrorMessage: "run job  - jar filename not set"}.SendResponse(w, r)
		return
	}

	go func() {
		lj.SubmitSparkJob(sr.Job, sr.Jar)
	}()

	responses.InProgressResponse{}.SendResponse(w, r)

}
