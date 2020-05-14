package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Config configuration

func init() {
	configFile, err := ioutil.ReadFile(fileName())

	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS Livy").
			Msgf("Cannot read configuration")
	}

	Config = configuration{}

	err = toml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS Livy").
			Msgf("Cannot unmarshall configuration file")
	}

	// Parse environment variables
	if err := env.Parse(&Config); err != nil {
		log.Fatal().
			Err(err).
			Str("service", "LFS Livy").
			Msgf("Cannot parse environment variables")
	}

	// setup logging

	switch Config.LogFormat {
	case "Text":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339})
	case "Json":
		break
	case "Terminal":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.RFC3339})
		break
	}

	switch Config.LogLevel {
	case "Trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "Info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "Debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "Warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "Error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "Fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

}

func fileName() string {
	runEnv := os.Getenv("CONFIG")

	if len(runEnv) == 0 {
		runEnv = "development"
	}

	filename := []string{"config.", runEnv, ".toml"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
