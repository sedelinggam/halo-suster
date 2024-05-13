package server

import (
	v1 "halo-suster/internal/delivery/http/v1"
	"halo-suster/package/database/postgresql"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	db        *sqlx.DB
	app       *echo.Echo
	validator *validator.Validate
}

func NewServer() *Server {
	e := echo.New()
	validate := validator.New()
	db := postgresql.New()

	return &Server{
		db:        db,
		app:       e,
		validator: validate,
	}
}

func (s *Server) Run() error {
	//Setup Middleware
	//Logger Middleware
	logger, _ := zap.NewProduction()
	s.app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	//Setup Router
	v1.Init(s.app, s.db, s.validator)

	return s.app.Start(":8080")
}
