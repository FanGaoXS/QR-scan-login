package qr

import (
	"context"
	"net/url"

	"fangaoxs.com/QR-scan-login/environment"
	"fangaoxs.com/QR-scan-login/internal/deps/qrcode"
	"fangaoxs.com/QR-scan-login/internal/domain/pincode"
	"fangaoxs.com/QR-scan-login/internal/infras/errors"
	"fangaoxs.com/QR-scan-login/internal/infras/logger"
)

type QR interface {
	GenerateQR(ctx context.Context) ([]byte, error)
	VerifyQR(ctx context.Context, code string) error
}

func New(
	env environment.Env,
	logger logger.Logger,
	pc pincode.PinCode,
) (QR, error) {
	// verify QR callback service url
	u, err := url.Parse(env.QRCallbackServiceURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath(env.QRCallbackPath)

	qc, err := qrcode.NewGenerator(env)
	if err != nil {
		return nil, err
	}
	return &qr{
		qrCallbackService: u.String(),
		pc:                pc,
		qc:                qc,
	}, nil
}

type qr struct {
	qrCallbackService string
	pc                pincode.PinCode
	qc                qrcode.Generator
}

const QrKey = "code"

func (q *qr) GenerateQR(ctx context.Context) ([]byte, error) {
	code, err := q.pc.Generate(ctx)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(q.qrCallbackService)
	values := u.Query()
	values.Add(QrKey, code)
	u.RawQuery = values.Encode()

	return q.qc.GenerateToBytes(ctx, u.String())
}

func (q *qr) VerifyQR(ctx context.Context, requestURI string) error {
	u, _ := url.ParseRequestURI(requestURI)
	values := u.Query()
	if !values.Has(QrKey) {
		return errors.Newf(errors.InvalidArgument, nil, "code not found")
	}

	code := values.Get(QrKey)
	return q.pc.Verify(ctx, code)
}
