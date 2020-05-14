package config

type configuration struct {
	LogFormat  string
	LogLevel   string
	LivyServer string `env:"LIVY_SERVER"`
	Service    ServiceConfiguration
}
