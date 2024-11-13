package example_test

import (
	"context"

	"github.com/ezraisw/test-gen/internal/example"
)

func runServiceUserDo(mocks []any) {
	mockService := mocks[0].(*example.MockService)
	u := example.NewServiceUser(mockService)
	_, _ = u.Do(context.Background())
}
