package pincode

import (
	"context"
	"sync"
	"time"

	"fangaoxs.com/QR-scan-login/internal/infras/errors"
)

type PinCode interface {
	Generate(ctx context.Context) (string, error)
	Verify(ctx context.Context, pin string) error
}

type pinPolicy struct {
	size       int
	expiration time.Duration
}

func New() (PinCode, error) {
	defaultPolicy := pinPolicy{
		size:       8,
		expiration: time.Second * 20,
	}

	return &pincode{
		Mutex:     sync.Mutex{},
		policy:    defaultPolicy,
		pins:      make(map[string]struct{}),
		expireAts: make(map[string]time.Time),
	}, nil
}

type pincode struct {
	sync.Mutex
	policy pinPolicy

	pins      map[string]struct{}
	expireAts map[string]time.Time
}

func (p *pincode) Generate(ctx context.Context) (string, error) {
	pin := randString(p.policy.size)

	if err := p.set(pin, p.policy.expiration); err != nil {
		return "", err
	}

	return pin, nil
}

func (p *pincode) Verify(ctx context.Context, pin string) error {
	return p.get(pin)
}

func (p *pincode) set(pin string, expiration time.Duration) error {
	p.Lock()
	defer p.Unlock()

	p.pins[pin] = struct{}{}
	p.expireAts[pin] = time.Now().Add(expiration)

	return nil
}

func (p *pincode) get(pin string) error {
	p.Lock()
	defer p.Unlock()

	expireAt, ok := p.expireAts[pin]
	if !ok {
		return errors.New(errors.InvalidArgument, nil, "invalid pin code or pin code has been used")
	}
	if time.Now().After(expireAt) {
		return errors.New(errors.DeadlineExceeded, nil, "pin code is expired")
	}

	return nil
}
