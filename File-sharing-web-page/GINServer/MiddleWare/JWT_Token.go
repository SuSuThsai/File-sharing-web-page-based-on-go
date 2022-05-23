package MiddleWare

import (
	//本地
	"errors"
	"net/http"
	"time"
	//第三方
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//目录
	"File-sharing-web-page/GINServer/GINconfig"
)

var GetJwtKey *GINconfig.JWT

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetJwtKey.JwtKey),
	}
}

type Claims struct {
	UserID uint `mapstructure:"userId" json:"user_id"`
	jwt.StandardClaims
}

//定义错误
var (
	TokenExpired     error = errors.New("token已过期,请重新登录")
	TokenNotValueYet error = errors.New("此token无效,请重新登录")
	TokenMalFormed   error = errors.New("token不正确,请重新登录")
	TokenInvalid     error = errors.New("token格式错误,请重新登录")
)

// CreatToken 生成TOKEN(J*JWT)
func (j *JWT) CreatToken(claims Claims) (string, error) {
	//expireTime := time.Now().Add(7 * 24 * time.Hour)
	//claims = Claims{
	//		UserID: Global.Global.User.ID,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expireTime.Unix(),
	//		IssuedAt:  time.Now().Unix(),
	//		Issuer:    "Yamada",
	//		Subject:   "u-token",
	//	},
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParserToken 验证TOKEN
func (j *JWT) ParserToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, nil, TokenMalFormed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, nil, TokenNotValueYet
			} else {
				return nil, nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claimsData, ok := token.Claims.(*Claims); ok && token.Valid {
			return nil, claimsData, nil
		}
		return nil, nil, TokenInvalid
	}

	return nil, nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreatToken(*claims)
	}
	return "", TokenInvalid
}

// JwtToken 中间键
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		//获取Authorization
		tokenString := context.Request.Header.Get("u-token")

		//判断token 的情况
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": "请先登录",
			})
			context.Abort()
			return
		}

		j := NewJWT()
		//解析Token
		_, claims, err := j.ParserToken(tokenString)
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": TokenExpired,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			context.Abort()
			return
		}
		context.Set("user_id", claims.Id)
		context.Set("claims", claims)
		context.Next()
	}
}
