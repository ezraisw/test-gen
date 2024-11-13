package example

import "context"

type Out struct {
	X string
	Y int
}

//go:generate mockgen -typed -source=service.go -destination=service_mock.go -package=example
type Service interface {
	A(ctx context.Context, in int) error
	B(ctx context.Context, in int) (*Out, error)
	C() int
}

type ServiceUser struct {
	service Service
}

func NewServiceUser(service Service) *ServiceUser {
	return &ServiceUser{
		service: service,
	}
}

func (u ServiceUser) Do(ctx context.Context) (int, error) {
	if err := u.service.A(ctx, 50); err != nil {
		return 0, err
	}

	b, err := u.service.B(ctx, 50)
	if err != nil {
		return 0, err
	}

	c := u.service.C()

	return b.Y + c, nil
}
