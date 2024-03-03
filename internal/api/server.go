package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"goods_project/internal/config"
	"goods_project/internal/service"
	"goods_project/internal/utils"
)

type Server struct {
	Config   *config.Config
	Service  *service.Service
	Log      *slog.Logger
	Validate *validator.Validate
	Ctx      context.Context
}

func NewServer(cfg *config.Config, service *service.Service, log *slog.Logger, ctx context.Context) *Server {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &Server{
		Config:   cfg,
		Service:  service,
		Log:      log,
		Validate: validate,
		Ctx:      ctx,
	}
}

func (s *Server) Start() {
	address := fmt.Sprintf("%s:%s", s.Config.App.Host, s.Config.App.Port)
	router := mux.NewRouter()

	router.HandleFunc("/good/", utils.MakeHTTPHandleFunc(s.CreateGood)).Methods(http.MethodPost)
	router.HandleFunc("/good/", utils.MakeHTTPHandleFunc(s.UpdateGood)).Methods(http.MethodPatch)
	router.HandleFunc("/good/", utils.MakeHTTPHandleFunc(s.GetGood)).Methods(http.MethodGet)
	router.HandleFunc("/good/reprioritize/", utils.MakeHTTPHandleFunc(s.Reprioritize)).Methods(http.MethodPatch)
	router.HandleFunc("/good/remove/", utils.MakeHTTPHandleFunc(s.RemoveGood)).Methods(http.MethodDelete)
	router.HandleFunc("/goods/list/", utils.MakeHTTPHandleFunc(s.ListGoods)).Methods(http.MethodGet)

	s.Log.Info("api is running", slog.String("host", s.Config.App.Host), slog.String("port", s.Config.App.Port))
	if err := http.ListenAndServe(address, router); err != nil {
		s.Log.Error("api is down", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
