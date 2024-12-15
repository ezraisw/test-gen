package analyzer_test

import (
	"testing"

	"github.com/ezraisw/test-gen/analyzer"
	"github.com/ezraisw/test-gen/internal/example"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAnalyzer(t *testing.T) {
	ctrl := gomock.NewController(nil)
	mock := example.NewMockSimple(ctrl)

	azr := analyzer.NewAnalyzer()
	azr.AttachTrap(mock, map[string][]any{
		"A": {"This is A"},
		"B": {"This is B"},
		"C": {"This is C", 50},
	})

	u := example.NewSimpleUser(mock)
	a, b, c1, c2 := u.Do()

	assert.Equal(t, "This is A", a)
	assert.Equal(t, "This is B", b)
	assert.Equal(t, "This is C", c1)
	assert.Equal(t, 50, c2)

	capturedCalls := azr.GetCapturedCalls()
	assert.Len(t, capturedCalls, 3)

	capturedCallStrs := make([]string, 0, len(capturedCalls))
	for _, capturedCall := range capturedCalls {
		capturedCallStrs = append(capturedCallStrs, capturedCall.String())
	}

	assert.Equal(t, []string{
		"github.com/ezraisw/test-gen/internal/example.MockSimple.A",
		"github.com/ezraisw/test-gen/internal/example.MockSimple.B",
		"github.com/ezraisw/test-gen/internal/example.MockSimple.C",
	}, capturedCallStrs)
}
