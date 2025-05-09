package logger

import (
	"context"
	"testing"
)

// nolint:staticcheck
func TestLogger(t *testing.T) {
	log := NewLogger(WithCtxFields([]string{"X-Request-ID", "X-Username"}))
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-ID", "X-Request-ID")
	ctx = context.WithValue(ctx, "X-Username", "X-Username")

	log.Info(ctx, "打印错误消息信息， err：%s, field:%s", "这里是err错误信息", "字段信息")
}
