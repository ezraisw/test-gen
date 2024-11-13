package analyzer

import (
	reflect "reflect"
	"sync"

	gomock "go.uber.org/mock/gomock"
)

var condAnyRv = reflect.ValueOf(gomock.Any())

type Analyzer struct {
	mu            sync.Mutex
	capturedCalls []*CallSignature
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

// Create a trap for using a gomock mock.
// Traps allows analyzer to detect called methods.
func (a *Analyzer) AttachTrap(r any, retsByMethodName map[string][]any) {
	rv := reflect.ValueOf(r)

	expectMethodRv := a.getExpectMethod(rv)

	rt := rv.Type()

	recorderRv := expectMethodRv.Call([]reflect.Value{})[0]
	if recorderRv.Kind() != reflect.Pointer {
		panic("return value from EXPECT is not a pointer to a possible recorder")
	}
	recorderRt := recorderRv.Type()

	for methodName := range retsByMethodName {
		_, ok := recorderRt.MethodByName(methodName)
		if !ok {
			panic("method name not found: " + methodName)
		}
	}

	for i := 0; i < recorderRt.NumMethod(); i++ {
		method := recorderRt.Method(i)
		if !method.IsExported() {
			continue
		}

		methodRv := recorderRv.MethodByName(method.Name)
		methodRt := methodRv.Type()

		if methodRt.NumOut() != 1 {
			panic("recorder method does not have exactly 1 return value")
		}

		argRvs := make([]reflect.Value, 0, methodRt.NumIn())
		for i := 0; i < methodRt.NumIn(); i++ {
			argRvs = append(argRvs, condAnyRv)
		}

		rets := retsByMethodName[method.Name]
		cs := &CallSignature{
			Type:       rt.Elem(),
			MethodName: method.Name,
			MethodType: methodRt,
			Returns:    rets,
		}

		retRvs := methodRv.Call(argRvs)
		a.attachCall(cs, retRvs[0], retsByMethodName[method.Name])
	}
}

func (a *Analyzer) GetCapturedCalls() []*CallSignature {
	return a.capturedCalls
}

func (a *Analyzer) getExpectMethod(mockRv reflect.Value) reflect.Value {
	if mockRv.Kind() != reflect.Ptr {
		panic("given value is not a pointer type to a possible gomock mock")
	}

	expectMethodRv := mockRv.MethodByName("EXPECT")
	if !expectMethodRv.IsValid() {
		panic("mock does not have EXPECT method")
	}
	// EXPECT is guaranteed a function kind.

	expectMethodRt := expectMethodRv.Type()
	if expectMethodRt.NumIn() != 0 {
		panic("EXPECT method does not have exactly 0 parameters")
	}

	if expectMethodRt.NumOut() != 1 {
		panic("EXPECT method does not have exactly 1 return value")
	}

	return expectMethodRv
}

func (a *Analyzer) attachCall(cs *CallSignature, callRv reflect.Value, rets []any) {
	if callRv.Kind() != reflect.Pointer {
		panic("return value from recorder method is not a pointer type to a possible gomock call")
	}

	doAndReturnMethodRv := callRv.MethodByName("DoAndReturn")
	doAndReturnMethodRt := doAndReturnMethodRv.Type()

	if doAndReturnMethodRt.NumIn() != 1 {
		panic("DoAndReturn does not have exactly 1 parameter")
	}

	if doAndReturnMethodRt.NumOut() != 1 {
		panic("DoAndReturn does not have exactly 1 return value")
	}

	realMethodRt := doAndReturnMethodRt.In(0)
	if realMethodRt.Kind() == reflect.Interface {
		panic("mock is generated without type information; cannot determine function type")
	}

	if rets != nil && realMethodRt.NumOut() != len(rets) {
		panic("number of return values does not match")
	}

	realMethodRv := reflect.MakeFunc(realMethodRt, func([]reflect.Value) []reflect.Value {
		a.mu.Lock()
		a.capturedCalls = append(a.capturedCalls, cs)
		a.mu.Unlock()

		retRvs := make([]reflect.Value, 0, realMethodRt.NumOut())
		if rets == nil {
			for i := 0; i < realMethodRt.NumOut(); i++ {
				retRvs = append(retRvs, reflect.New(realMethodRt.Out(i)).Elem())
			}
		} else {
			for i, ret := range rets {
				retRv := reflect.ValueOf(ret)
				if retRv.Kind() == reflect.Invalid {
					retRv = reflect.New(realMethodRt.Out(i)).Elem()
				}
				retRvs = append(retRvs, retRv)
			}
		}

		return retRvs
	})

	retRvs := doAndReturnMethodRv.Call([]reflect.Value{realMethodRv})
	callRv = retRvs[0]

	anyTimesMethodRv := callRv.MethodByName("AnyTimes")
	anyTimesMethodRt := anyTimesMethodRv.Type()

	if anyTimesMethodRt.NumIn() != 0 {
		panic("AnyTimes does not have exactly 0 parameters")
	}

	if anyTimesMethodRt.NumOut() != 1 {
		panic("AnyTimes does not have exactly 1 return value")
	}

	anyTimesMethodRv.Call([]reflect.Value{})
}
