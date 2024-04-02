package rest

import (
	"fmt"
	"net/http"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/domain/qr"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"

	"github.com/gin-gonic/gin"
)

func newHandlers(
	env environment.Env,
	logger logger.Logger,
	qr qr.QR,
) (handlers, error) {
	return handlers{
		qr: qr,
	}, nil
}

type handlers struct {
	qr qr.QR
}

func (h *handlers) GenerateQR() gin.HandlerFunc {
	return func(c *gin.Context) {
		// GET
		ctx := c.Request.Context()

		png, err := h.qr.GenerateQR(ctx)
		if err != nil {
			WrapGinError(c, err)
			return
		}

		c.Data(http.StatusOK, "image/png", png)
	}
}

func (h *handlers) VerifyQR() gin.HandlerFunc {
	return func(c *gin.Context) {
		// GET
		ctx := c.Request.Context()

		remoteIP := c.RemoteIP()
		clientIP := c.ClientIP()
		agent := c.Request.UserAgent()
		fmt.Printf("remoteIP = %s, clientIP = %s, user-agent = %s\n", remoteIP, clientIP, agent)

		if err := h.qr.VerifyQR(ctx, c.Request.RequestURI); err != nil {
			WrapGinError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		})
	}
}
