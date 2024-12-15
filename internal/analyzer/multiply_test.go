package analyzer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ezraisw/test-gen/internal/analyzer"
	"github.com/ezraisw/test-gen/internal/example"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMultiply(t *testing.T) {
	mockErrA1 := errors.New("mock error A 1")
	mockErrA2 := errors.New("mock error A 2")
	mockErrB := errors.New("mock error B")

	resOuts := make([]int, 0)
	resErrs := make([]error, 0)
	res := analyzer.Multiply([]*analyzer.MockConfig{
		{
			New: func(ctrl *gomock.Controller) any { return example.NewMockService(ctrl) },
			Methods: []*analyzer.MockMethod{
				{Name: "A", Returns: analyzer.Vary{
					analyzer.Stop{mockErrA1},
					analyzer.Stop{mockErrA2},
					analyzer.Pass{nil},
				}},
				{Name: "B", Returns: analyzer.Vary{
					analyzer.Stop{nil, mockErrB},
					analyzer.Pass{&example.Out{X: "This is X", Y: 50}, nil},
				}},
				{Name: "C", Returns: analyzer.Pass{35}},
			},
		},
	}, func(mocks []any) {
		mock := mocks[0].(*example.MockService)
		u := example.NewServiceUser(mock)
		out, err := u.Do(context.Background())

		resOuts = append(resOuts, out)
		resErrs = append(resErrs, err)
	})

	assert.Equal(t, []int{0, 0, 0, 85}, resOuts)
	assert.Equal(t, []error{mockErrA1, mockErrA2, mockErrB, nil}, resErrs)

	expectedRes := [][]string{
		{"github.com/ezraisw/test-gen/internal/example.MockService.A"},
		{"github.com/ezraisw/test-gen/internal/example.MockService.A"},
		{"github.com/ezraisw/test-gen/internal/example.MockService.A", "github.com/ezraisw/test-gen/internal/example.MockService.B"},
		{"github.com/ezraisw/test-gen/internal/example.MockService.A", "github.com/ezraisw/test-gen/internal/example.MockService.B", "github.com/ezraisw/test-gen/internal/example.MockService.C"},
	}

	assert.Len(t, res, len(expectedRes))

	for i, capturedCalls := range res {
		capturedCallStrs := make([]string, 0, len(capturedCalls))
		for _, capturedCall := range capturedCalls {
			capturedCallStrs = append(capturedCallStrs, capturedCall.String())
		}

		assert.Equal(t, expectedRes[i], capturedCallStrs)
	}
}
