// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initServer(env environment.Env, logger2 logger.Logger, httpServer *gin.Engine) (*Server, error) {
	server, err := newServer(env, logger2, httpServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
