package qrcode

import (
	"context"
	"encoding/base64"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/infras/errors"

	goqrcode "github.com/skip2/go-qrcode"
)

type Generator interface {
	GenerateToBase64String(ctx context.Context, content string) (string, error)
	GenerateToBytes(ctx context.Context, content string) ([]byte, error)
}

func NewGenerator(env environment.Env) (Generator, error) {
	return &generator{}, nil
}

type generator struct{}

func (g *generator) GenerateToBase64String(ctx context.Context, content string) (string, error) {
	png, err := goqrcode.Encode(content, goqrcode.Highest, 256)
	if err != nil {
		return "", errors.Newf(errors.Internal, nil, "generate qrcode failed")
	}

	base64String := base64.StdEncoding.EncodeToString(png)
	return "data:image/png;base64," + base64String, nil
}

func (g *generator) GenerateToBytes(ctx context.Context, content string) ([]byte, error) {
	png, err := goqrcode.Encode(content, goqrcode.Highest, 256)
	if err != nil {
		return nil, errors.Newf(errors.Internal, nil, "generate qrcode failed")
	}

	return png, nil
}
