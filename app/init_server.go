package app

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"tracker/internal/application/service"
	"tracker/internal/application/translator"
	board2 "tracker/internal/interface/http/board"
	project2 "tracker/internal/interface/http/project"
	workspace2 "tracker/internal/interface/http/workspace"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	return &Server{e: echo.New()}
}

func (s *Server) InitServer() {
	s.connectMiddlewares()
	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	s.connectApiRoutes()
	data, err := json.MarshalIndent(s.e.Routes(), "", "  ")
	if err != nil {
		log.Err(err).Msg("err")
	}
	log.Info().Bytes("routes", data).Msg("routes")
	log.Err(s.e.Start(":6030")).Msg("error to start server")
}

func (s *Server) connectApiRoutes() {
	gr := s.e.Group("api/v1/")
	gr.Use(echojwt.WithConfig(echojwt.Config{
		// ...
		SigningKey: []byte("secret"),
		// ...
	}))

	validate := validator.New(validator.WithRequiredStructEnabled())
	errTranslator := translator.NewTranslator(validate)

	wsService := service.NewWorkspaceService(GetDb())
	workspaceGr := gr.Group("workspace")
	workspaceGr.POST("", workspace2.NewCreateHandler(wsService, errTranslator, validate).Handle)

	projectService := service.NewProjectService(GetDb())
	projectGr := gr.Group("project")
	projectGr.POST("", project2.NewCreateHandler(projectService).Handle)

	boardService := service.NewBoardService(GetDb())
	boardGr := gr.Group("board")
	boardGr.POST("", board2.NewCreateHandler(boardService).Handle)
}

func (s *Server) connectMiddlewares() {
	s.e.Use(middleware.RemoveTrailingSlash())
	s.e.Use(middleware.RequestID())
	s.e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Info().
			Str("request_body", string(reqBody)).
			Str("response_body", string(resBody))
	}))
	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogError:     true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Err(v.Error).
					Str("Method", v.Method).
					Str("URI", v.URI).
					Int("status", v.Status).
					Time("start_time", v.StartTime).
					Str("request_id", v.RequestID).
					Msg("request")
			} else {
				log.Info().
					Str("Method", v.Method).
					Str("URI", v.URI).
					Int("status", v.Status).
					Time("start_time", v.StartTime).
					Str("request_id", v.RequestID).
					Msg("request")
			}

			return nil
		},
	}))
	s.e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(context echo.Context, reqBody []byte, resBody []byte) {
			log.Info().
				Str("request_body", string(reqBody)).
				Str("response_body", string(resBody))
		},
	}))
}
