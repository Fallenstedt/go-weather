package util_test

import (
	"context"
	"testing"

	"github.com/Fallenstedt/weather/common/util"
)

func TestGetFlagsFromContext(t *testing.T) {
	flags := util.Flags{
		Detail: 1,
	}

	ctx := context.WithValue(context.Background(), util.ContextKeyFlags, flags)

	result := util.GetFlagsFromContext(ctx)

	if result.Detail != flags.Detail {
		t.Errorf("got %v, expected %v", result, flags)
	}
}