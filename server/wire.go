//go:build wireinject
// +build wireinject

package server

import (
	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/domain/pincode"
	"fangaoxs.com/QR-scan-login/internal/domain/qr"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func initServer(env environment.Env, logger logger.Logger, httpServer *gin.Engine) (*Server, error) {
	panic(wire.Build(
		pincode.New,
		qr.New,
		newServer,
	))
}
