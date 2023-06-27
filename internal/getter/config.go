package getter

import "fmt"

type Config struct {
	ApiKey string
	Host   string
}

func NewConfig(apiKey string, host string) (*Config, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("api key is empty")
	}
	if host == "" {
		return nil, fmt.Errorf("host is empty")
	}
	return &Config{
		ApiKey: apiKey,
		Host:   host,
	}, nil
}
