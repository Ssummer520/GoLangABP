package Authentication



import (
//"awesomeProject/utils"
"github.com/dgrijalva/jwt-go"
"github.com/gin-gonic/gin"
"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}
var errObject interface{ }
var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//该路由下不校验token
	noVerify = []interface{}{"/login", "/ping"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		errObject=err
		panic(errObject)
	}
	return sign
}

// JwtVerify 验证token
func JwtVerify(c *gin.Context) {
	return
	//过滤是否验证token
	//if utils.IsContainArr(noVerify, c.Request.RequestURI) {
	//	return
	//}
	token := c.GetHeader("token")
	if token == "" {
		errObject = "token not exist !"
		panic(errObject)
	}
	//验证token，并存储在请求中
	c.Set("user", parseToken(token))
}

// 解析Token
func parseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		var v1 interface{ }
		v1 = err
		panic(v1)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		errObject = "token is valid"
		panic(errObject)
	}
	return claims
}

// Refresh 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		errObject = err
		panic(errObject)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		errObject= "token is valid"
		panic(errObject)
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}