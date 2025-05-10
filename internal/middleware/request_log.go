package middleware

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-study/pkg/app/core"
	"gin-study/pkg/logger"
	"gin-study/pkg/utils"
)

// RequestLog 记录请求的日志信息.
func RequestLog(logger *logger.ZapLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置traceId, 用来进行链路追踪
		traceId := ctx.GetHeader(core.HeaderKeyTraceID)
		if traceId == "" {
			traceId = utils.RandStr(10)
			ctx.Header(core.HeaderKeyTraceID, traceId)
		}
		ctx.Set(core.HeaderKeyTraceID, traceId)

		// 打印请求日志
		timer := time.Now()
		rowData, _ := ctx.GetRawData()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(rowData))
		logger.Info(
			ctx,
			"%-6s %-10s %-25s req:\n%s\n",
			ctx.Request.Method,
			filepath.Base(ctx.HandlerName()),
			ctx.Request.RequestURI,
			strings.ReplaceAll(string(rowData), "\r\n", ""),
		)

		ctx.Next()

		logger.Info(ctx, "耗时：%.2fs", time.Since(timer).Seconds())
	}
}
