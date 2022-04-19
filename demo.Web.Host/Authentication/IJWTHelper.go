package Authentication

import "github.com/gin-gonic/gin"

// IJWTHelper JWT相关接口
type IJWTHelper interface {
	// Refresh  刷新token
	Refresh(tokenString string) string
	// 解析Token
	parseToken(tokenString string) *UserClaims
	// JwtVerify 验证token
	JwtVerify(c *gin.Context)
	// GenerateToken 生成Token
	GenerateToken(claims *UserClaims) string
}
