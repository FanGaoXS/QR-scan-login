package qrcode

import (
	"context"
	"testing"

	"fangaoxs.com/QR-scan-login/environment"
)

func TestQRCode(t *testing.T) {
	ctx := context.Background()

	env, err := environment.Get()
	if err != nil {
		t.Fatal(err)
	}

	qr, err := New(env)
	if err != nil {
		t.Fatal(err)
	}

	code, err := qr.Generate(ctx, "test")
	if err != nil {
		t.Error(err)
	}

	println(code)
}
