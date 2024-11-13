package analyzer

import reflect "reflect"

type CallSignature struct {
	Type       reflect.Type
	MethodName string
	MethodType reflect.Type
	Returns    []any
}

func (c CallSignature) String() string {
	return c.Type.PkgPath() + "." + c.Type.Name() + "." + c.MethodName
}
