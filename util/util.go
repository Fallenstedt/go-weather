package util

import (
	"context"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

type Flags struct {
	Detail int
}

var (
	ContextKeyFlags = contextKey("flags")
)

func GetFlagsFromContext(ctx context.Context) Flags {
	flags, ok := ctx.Value(ContextKeyFlags).(Flags)
	if !ok {
		panic("Failed to convert flags type")
	}
	return flags
}
