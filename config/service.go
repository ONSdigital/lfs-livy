package config

type ServiceConfiguration struct {
	ListenAddress string `env:"LISTEN_ADDRESS"`
	ReadTimeout   string
	WriteTimeout  string
}
