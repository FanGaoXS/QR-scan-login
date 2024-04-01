package pincode

import (
	"context"
	"testing"
	"time"
)

func TestPinCode(t *testing.T) {
	ctx := context.Background()
	p, _ := New()
	pin, _ := p.Generate(ctx)

	time.Sleep(time.Second * 5)
	err := p.Verify(ctx, pin)
	if err != nil {
		t.Errorf("should not err")
	}

	time.Sleep(time.Second * 5)
	err = p.Verify(ctx, pin)
	if err == nil {
		t.Errorf("should err")
	}

	pin, _ = p.Generate(ctx)
	err = p.Verify(ctx, pin)
	if err != nil {
		t.Errorf("should not err")
	}
}
