package analyzer

import "go.uber.org/mock/gomock"

type MockConfig struct {
	New     func(ctrl *gomock.Controller) any
	Methods []*MockMethod
}

type MockMethod struct {
	Name    string
	Returns VaryLike
}

func Multiply(mockCfgs []*MockConfig, run func(mocks []any)) [][]*CallSignature {
	dt := newVaryDecisionTree(mockCfgs, run)
	dt.traverse(0)

	return dt.allCapturedCalls
}

type varyPathSelection struct {
	vary Vary

	// Selected mock method's return value.
	varyIdx int
}

type varyDecisionTree struct {
	// Configuration.
	mockCfgs []*MockConfig
	run      func(mocks []any)

	// Recursion information.
	pathSelections []*varyPathSelection
	toPsIdx        map[[2]int]int

	// Result.
	allCapturedCalls [][]*CallSignature
}

func newVaryDecisionTree(mockCfgs []*MockConfig, run func(mocks []any)) *varyDecisionTree {
	pathSelections := make([]*varyPathSelection, 0)
	toPsIdx := make(map[[2]int]int)

	for mockCfgIdx, mockCfg := range mockCfgs {
		for methodIdx, method := range mockCfg.Methods {
			key := [2]int{mockCfgIdx, methodIdx}
			toPsIdx[key] = len(pathSelections) // Index at that time.
			pathSelections = append(pathSelections, &varyPathSelection{
				vary: method.Returns.asVary(),

				// Set to -1 to indicate no selection.
				varyIdx: -1,
			})
		}
	}

	return &varyDecisionTree{
		mockCfgs: mockCfgs,
		run:      run,

		pathSelections: pathSelections,
		toPsIdx:        toPsIdx,

		allCapturedCalls: make([][]*CallSignature, 0),
	}
}

func (dt *varyDecisionTree) traverse(psIdx int) {
	if len(dt.pathSelections) == psIdx {
		dt.pathFinish()
		return
	}

	currPs := dt.pathSelections[psIdx]

	for varyIdx := 0; varyIdx < len(currPs.vary); varyIdx++ {
		nextPsIdx := psIdx + 1

		if !currPs.vary[varyIdx].passable() {
			// Set nextPsIdx to short circuit the recursion.
			nextPsIdx = len(dt.pathSelections)
		}

		// Restore to previous state after recursion for safety.
		prevVaryIdx := currPs.varyIdx
		currPs.varyIdx = varyIdx
		dt.traverse(nextPsIdx)
		currPs.varyIdx = prevVaryIdx
	}
}

func (dt *varyDecisionTree) pathFinish() {
	ctrl := gomock.NewController(nil)
	azr := NewAnalyzer()

	mocks := make([]any, 0, len(dt.mockCfgs))
	for mockCfgIdx, mockCfg := range dt.mockCfgs {
		mock := mockCfg.New(ctrl)
		mocks = append(mocks, mock)

		retsByMethodName := make(map[string][]any)
		for methodIdx, method := range mockCfg.Methods {
			psIdx := dt.toPsIdx[[2]int{mockCfgIdx, methodIdx}]
			ps := dt.pathSelections[psIdx]

			// Skip defining a return value for paths that are not selected.
			if ps.varyIdx == -1 {
				continue
			}

			retsByMethodName[method.Name] = method.Returns.getAt(ps.varyIdx)
		}

		azr.AttachTrap(mock, retsByMethodName)
	}

	dt.run(mocks)

	dt.allCapturedCalls = append(dt.allCapturedCalls, azr.GetCapturedCalls())
}
