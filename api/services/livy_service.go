package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ONSDigital/lfs-livy/config"
	"github.com/rs/zerolog/log"
	"net/http"
)

type LivyRequestService struct {
}

type StartLivyJob struct {
	Name      string `json:"name"`
	File      string `json:"file"`
	ClassName string `json:"className"`
}

func (lr LivyRequestService) SubmitSparkJob(jobName string, jarFile string) {

	livySession := StartLivyJob{
		Name:      jobName, // this has to be unique i.e. no recycled names
		File:      jarFile,
		ClassName: "uk.gov.ons.lfs.LFSMonthly",
	}

	jsonData, err := json.Marshal(livySession)
	if err != nil {
		log.Error().
			Err(err).
			Msg("cannot create a livy session - marshall json failed")
		return
	}

	var url = getLivyServer() + "/batches"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().
			Err(err).
			Msg("cannot create a livy session - failed on NewRequest")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-By", "LivyServer")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Error().
			Err(err).
			Msg("cannot create a livy session - failed on send")
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

func getLivyServer() string { return config.Config.LivyServer }
