package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/joho/godotenv"
	"github.com/prometheus/common/model"
)

func NewLokiLogger() (*loki.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	endpoint := os.Getenv("LOKI_ENDPOINT")
	config, err := loki.NewDefaultConfig(endpoint)
	if err != nil {
		return nil, err
	}

	client, err := loki.New(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func SendLoki(client *loki.Client, labels model.LabelSet) error {
	err := client.Handle(labels, time.Now(), "")

	if err != nil {
		return fmt.Errorf("")
	}

	return nil
}
