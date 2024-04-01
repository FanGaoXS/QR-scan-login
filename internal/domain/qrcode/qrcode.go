package qrcode

import (
	"context"
	"encoding/base64"
	"net/url"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/infras/errors"

	goqrcode "github.com/skip2/go-qrcode"
)

type QRCode interface {
	Generate(ctx context.Context, code string) (string, error)
}

func New(env environment.Env) (QRCode, error) {
	u, err := url.Parse(env.Domain)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath("qrcode")

	return &qrcode{
		domain: u,
	}, nil
}

type qrcode struct {
	domain *url.URL
}

func (q *qrcode) Generate(ctx context.Context, code string) (string, error) {
	// {domain}/qrcode/{code}
	u := q.domain.JoinPath(code)
	content := u.String()

	png, err := goqrcode.Encode(content, goqrcode.Highest, 256)
	if err != nil {
		return "", errors.Newf(errors.Internal, nil, "generate qrcode failed")
	}

	base64String := base64.StdEncoding.EncodeToString(png)
	return "data:image/png;base64," + base64String, nil
}
