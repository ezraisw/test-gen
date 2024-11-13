package example

//go:generate mockgen -typed -source=simple.go -destination=simple_mock.go -package=example
type Simple interface {
	A() string
	B(a int) string
	C(b, c int) (string, int)
}

type SimpleUser struct {
	simple Simple
}

func NewSimpleUser(simple Simple) *SimpleUser {
	return &SimpleUser{
		simple: simple,
	}
}

func (u SimpleUser) Do() (string, string, string, int) {
	a := u.simple.A()
	b := u.simple.B(30)
	c1, c2 := u.simple.C(60, 90)

	return a, b, c1, c2
}
