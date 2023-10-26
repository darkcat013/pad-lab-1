package utils

import (
	"context"
	"time"
)

func GetDeadlineContext() (context.Context, context.CancelFunc) {
	return context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
}
