package main

import (
	"context"
	"fmt"
	"time"

	"github.com/paul-milne/zap-loki"
	"go.uber.org/zap"
)

func main() {
	v := zaploki.New(context.Background(), zaploki.Config{
		Url:          "http://localhost:3100",
		BatchMaxSize: 100,
		BatchMaxWait: 10 * time.Second,
		Labels:       map[string]string{"app": "test", "env": "dev"},
	})
	logger, err := v.WithCreateLogger(zap.NewProductionConfig())
	fmt.Println("WithCreateLogger", err)

	for {
		logger.Info("failed to fetch URL",
			// Structured context as strongly typed Field values.
			zap.String("url", "https.//test.com"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		defer logger.Sync()
	}
}
