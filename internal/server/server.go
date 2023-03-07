package server

import (
	"GolangCurrencyMS/config"
	"GolangCurrencyMS/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type server struct {
	controllers []handler.Controller
	config      *config.Config
}

func NewServer(cfg *config.Config, ctrls ...handler.Controller) Server {
	return &server{
		controllers: ctrls,
		config:      cfg,
	}
}

func (s *server) Run() error {
	// gin.SetMode(gin.ReleaseMode)

	router, err := s.buildRouter()
	if err != nil {
		return err
	}

	log.Printf("%d\n", s.config.App.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.App.Port),
		Handler: router,
	}

	startupFailed := make(chan interface{}, 1)
	defer close(startupFailed)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			startupFailed <- true
		}
	}()

	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)

	select {
	case <-quit:
		log.Println("shutdown web server ...")
		return nil
	case <-startupFailed:
		return fmt.Errorf("error on server start")
	}
}

func (s *server) buildRouter() (*gin.Engine, error) {
	router := gin.Default()

	for _, ctrl := range s.controllers {
		method := ctrl.Method()

		switch method {
		case http.MethodGet:
			router.GET(ctrl.Route(), ctrl.Handler())

		default:
			return nil, fmt.Errorf("Cannot create HTTP handler of method %s", method)
		}
	}
	return router, nil
}
