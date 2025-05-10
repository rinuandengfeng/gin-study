package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gin-study/pkg/app/core"
	"gin-study/pkg/logger"
)

// Recovery 将错误信息和panic信息输出到日志信息中，返回系统错误信息.
func Recovery(logger *logger.ZapLogger) gin.HandlerFunc {
	buffer := new(strings.Builder)

	return gin.CustomRecoveryWithWriter(buffer, func(c *gin.Context, recovered any) {
		logger.Error(c, "服务异常错误：%s, 打印panic信息\n%s", recovered, buffer.String())

		core.NewApp(c).Error(core.ErrSystem)
	})
}
