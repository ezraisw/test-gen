package analyzer

import (
	"go.uber.org/mock/gomock"
)

type Vary []any

type Cutoff []any

type Method struct {
	Name    string
	Returns any
}

func (m Method) getReturns(i int) []any {
	if vary, ok := m.Returns.(Vary); ok {
		if cutoff, ok := vary[i].(Cutoff); ok {
			return []any(cutoff)
		}
		if rets, ok := vary[i].([]any); ok {
			return rets
		}
		return nil
	}
	if cutoff, ok := m.Returns.(Cutoff); ok {
		return []any(cutoff)
	}
	if rets, ok := m.Returns.([]any); ok {
		return rets
	}
	return nil
}

type MockConfig struct {
	New     func(ctrl *gomock.Controller) any
	Methods []*Method
}

type varyDecision struct {
	mockCfgIdx int
	methodIdx  int
	vary       Vary
	idx        int
}

func Multiply(mockCfgs []*MockConfig, run func(mocks []any)) [][]*CallSignature {
	varyDecisions := make([]*varyDecision, 0)

	for mockCfgIdx, mockCfg := range mockCfgs {
		for methodIdx, method := range mockCfg.Methods {
			vary, ok := method.Returns.(Vary)
			if !ok {
				continue
			}

			varyDecisions = append(varyDecisions, &varyDecision{
				mockCfgIdx: mockCfgIdx,
				methodIdx:  methodIdx,
				vary:       vary,
				idx:        -1,
			})
		}
	}

	res := make([][]*CallSignature, 0)
	recurse(&res, mockCfgs, run, varyDecisions, 0)
	return res
}

func recurse(res *[][]*CallSignature, mockCfgs []*MockConfig, run func(mocks []any), varyDecisions []*varyDecision, i int) {
	if len(varyDecisions) == i {
		m := make(map[[2]int]int)

		for _, varyDecision := range varyDecisions {
			key := [2]int{varyDecision.mockCfgIdx, varyDecision.methodIdx}
			m[key] = varyDecision.idx
		}

		ctrl := gomock.NewController(nil)
		azr := NewAnalyzer()

		mocks := make([]any, 0, len(mockCfgs))
		for mockCfgIdx, mockCfg := range mockCfgs {
			mock := mockCfg.New(ctrl)
			mocks = append(mocks, mock)

			retsByMethodName := make(map[string][]any)
			for methodIdx, method := range mockCfg.Methods {
				varyIdx := m[[2]int{mockCfgIdx, methodIdx}]
				if varyIdx == -1 {
					continue
				}
				retsByMethodName[method.Name] = method.getReturns(varyIdx)
			}

			azr.AttachTrap(mock, retsByMethodName)
		}

		run(mocks)

		*res = append(*res, azr.GetCapturedCalls())
		return
	}

	for varyIdx := 0; varyIdx < len(varyDecisions[i].vary); varyIdx++ {
		nextI := i + 1
		if _, ok := varyDecisions[i].vary[varyIdx].(Cutoff); ok {
			nextI = len(varyDecisions)
		}

		prevVaryIdx := varyDecisions[i].idx

		varyDecisions[i].idx = varyIdx
		recurse(res, mockCfgs, run, varyDecisions, nextI)
		varyDecisions[i].idx = prevVaryIdx
	}
}
