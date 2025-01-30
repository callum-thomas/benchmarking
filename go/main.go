package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type HealthcheckResponse struct {
	Status string `json:"status"`
}

func main() {
	logger, _ := zap.NewDevelopment()

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogUserAgent: true,
		LogMethod:    true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("method", v.Method),
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.Int64("latency_ms", v.Latency.Milliseconds()),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("host", v.Host),
				zap.String("user_agent", v.UserAgent),
			)
			return nil
		},
	}))

	e.GET("/simpleHealthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "service healthy")
	})
	e.GET("/healthcheck", func(c echo.Context) error {
		resp := HealthcheckResponse{Status: "healthy"}
		return c.JSON(http.StatusOK, resp)
	})

	logger.Info("starting service")

	if err := e.Start(":8080"); err != nil {
		logger.Fatal("failure in server", zap.Error(err))
	}
}
