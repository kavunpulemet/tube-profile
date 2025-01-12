package api

import (
	"net/http"
	"time"
	"tube-profile/internal/api/handler"
	"tube-profile/internal/config"
	"tube-profile/internal/service"

	"github.com/gorilla/mux"

	"tube-profile/internal/api/middleware"
	"tube-profile/internal/utils"
)

const (
	maxHeaderBytes = 1 << 20
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(ctx utils.MyContext, config config.Config) *Server {
	router := mux.NewRouter()

	recoveredRouter := middleware.RecoveryMiddleware(ctx, router)

	authenticatedRouter := middleware.UserIdentity(ctx, recoveredRouter)

	return &Server{
		httpServer: &http.Server{
			Addr:           config.ServerPort,
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        authenticatedRouter,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx utils.MyContext) error {
	return s.httpServer.Shutdown(ctx.Ctx)
}

func (s *Server) HandleAuth(ctx utils.MyContext, service service.ProfileService) {
	s.router.HandleFunc("/api/profile/", handler.CreateProfile(ctx, service)).Methods(http.MethodPost)
	s.router.HandleFunc("/api/profile/", handler.GetProfile(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/api/profile/", handler.UpdateProfile(ctx, service)).Methods(http.MethodPut)
}
