package rest

import (
	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"

	"github.com/gin-gonic/gin"
)

func newHandlers(env environment.Env, logger logger.Logger) (handlers, error) {
	return handlers{}, nil
}

type handlers struct{}

func (h *handlers) GenerateCode() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *handlers) VerifyCode() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
