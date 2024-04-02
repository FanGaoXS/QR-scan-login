package rest

import (
	"net/http"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/domain/qr"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"

	"github.com/gin-gonic/gin"
)

func New(
	env environment.Env,
	logger logger.Logger,
	router *gin.Engine,
	qr qr.QR,
) (*Server, error) {
	qrCallbackPath := env.QRCallbackPath
	hdls, err := newHandlers(env, logger, qr)
	if err != nil {
		return nil, err
	}

	router.GET("generateQR", hdls.GenerateQR())
	router.GET(qrCallbackPath, hdls.VerifyQR())

	s := &http.Server{
		Addr:    env.RestListenAddr,
		Handler: router,
	}
	return &Server{
		server: s,
	}, nil
}

type Server struct {
	server *http.Server
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
