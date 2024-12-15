//go:build testgen
// +build testgen

package example

import (
	"context"
	"errors"

	testgen "github.com/ezraisw/test-gen"
	"github.com/ezraisw/test-gen/analyzer"
	"go.uber.org/mock/gomock"
)

//go:generate go run -tags testgen ./testgenrun
func Generate() {
	testgen.Generate("example_test", "service_cov_test.go", []*testgen.Test{
		{
			Name: "TestCoverage_ServiceUser_Do",
			MockConfigs: []*analyzer.MockConfig{
				{
					New: func(ctrl *gomock.Controller) any { return NewMockService(ctrl) },
					Methods: []*analyzer.Method{
						{Name: "A", Returns: analyzer.Vary{
							analyzer.Cutoff{errors.New("mock error A 1")},
							analyzer.Cutoff{errors.New("mock error A 2")},
							[]any{nil},
						}},
						{Name: "B", Returns: analyzer.Vary{
							analyzer.Cutoff{nil, errors.New("mock error B")},
							[]any{&Out{X: "This is X", Y: 50}, nil},
						}},
						{Name: "C", Returns: []any{35}},
					},
				},
			},
			Run: func(mocks []any) {
				mock := mocks[0].(*MockService)
				u := NewServiceUser(mock)
				_, _ = u.Do(context.Background())
			},
			TestRun: "runServiceUserDo",
		},
	})
}
