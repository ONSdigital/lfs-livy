package main

import (
	livy "github.com/ONSDigital/lfs-livy/api/handlers"
	"github.com/ONSDigital/lfs-livy/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	router := mux.NewRouter()

	log.Info().
		Str("startTime", time.Now().String()).
		Msg("LFS Livy: Starting up")

	// we'll allow anything for now. May need or want to restrict this to just the UI when we know its endpoint
	origins := []string{"*"}
	var cors = handlers.AllowedOrigins(origins)

	handlers.CORS(cors)(router)

	jobHandler := livy.NewLivyJobHandler()

	router.HandleFunc("/submit", jobHandler.RunJobHandler).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         getListenAddress(),
		WriteTimeout: getWriteTimeout(),
		ReadTimeout:  getReadtimeout(),
	}

	log.Info().
		Str("listenAddress", getListenAddress()).
		Str("writeTimeout", getWriteTimeout().String()).
		Str("readTimeout", getReadtimeout().String()).
		Msg("LFS Livy: Waiting for requests")

	err := srv.ListenAndServe()
	log.Fatal().
		Err(err).
		Str("service", "LFS").
		Msgf("ListenAndServe failed")
}

func getListenAddress() string { return config.Config.Service.ListenAddress }

func getWriteTimeout() time.Duration {
	writeTimeout, err := time.ParseDuration(config.Config.Service.WriteTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS Livy").
			Msgf("writeTimeout configuration error")
	}
	return writeTimeout
}

func getReadtimeout() time.Duration {
	readTimeout, err := time.ParseDuration(config.Config.Service.ReadTimeout)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS Livy")
	}
	return readTimeout
}
