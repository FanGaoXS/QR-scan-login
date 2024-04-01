package server

import (
	"context"
	"golang.org/x/sync/errgroup"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"
	"fangaoxs.com/QR-scan-login/server/rest"

	"github.com/gin-gonic/gin"
)

func New(env environment.Env, logger logger.Logger) (*Server, error) {
	httpServer := gin.New()
	gin.ForceConsoleColor()
	httpServer.Use(gin.Logger())

	server, err := initServer(env, logger, httpServer)
	if err != nil {
		return nil, err
	}

	return server, nil
}

type Server struct {
	env    environment.Env
	logger logger.Logger

	restServer *rest.Server
}

func newServer(
	env environment.Env,
	logger logger.Logger,
	httpServer *gin.Engine,
) (*Server, error) {
	restServer, err := rest.New(env, logger, httpServer)
	if err != nil {
		return nil, err
	}

	return &Server{
		env:        env,
		logger:     logger,
		restServer: restServer,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		s.logger.Infof("rest server listen on %s", s.env.RestListenAddr)
		err := s.restServer.ListenAndServe()
		if err != nil {
			return err
		}
		s.logger.Info("rest server stopped")
		return nil
	})

	go func() {
		select {
		case <-ctx.Done():
			s.close()
		}
	}()

	defer s.close()

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (s *Server) close() error {
	s.restServer.Close()
	return nil
}
