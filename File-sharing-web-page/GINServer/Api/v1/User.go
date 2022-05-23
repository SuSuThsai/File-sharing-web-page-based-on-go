package v1

import (
	//本地
	"log"
	"net/http"
	"time"
	//第三方
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//目录
	"File-sharing-web-page/GINServer/GINModels/User"
	"File-sharing-web-page/GINServer/GINconfig"
	"File-sharing-web-page/GINServer/MiddleWare"
	"File-sharing-web-page/GINServer/Validate"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data GINconfig.UserInfo
	err := c.ShouldBind(&data)
	if err != nil {
		log.Fatal(err.Error())
	}
	msg, validCode := Validate.Validate(&data)
	if !validCode {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  validCode,
			"message": msg,
		})
		c.Abort()
		return
	}
	code := User.CheckUser(data.Username)
	if code == http.StatusOK {
		User.CreatUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": "创建用户成功",
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"status":  code,
			"data":    data,
			"message": "创建用户失败",
		})
	}
}

// Login 后台登录
func Login(c *gin.Context) {
	var data GINconfig.UserInfo
	err := c.ShouldBind(&data)
	if err != nil {
		log.Fatal(err.Error())
	}
	userinfo, code := User.ValidateLogin(data.Username, data.Password)
	if code == http.StatusOK {
		setToken(c, userinfo)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"token":   nil,
			"message": "账号或密码错误",
		})
	}

}

//token生成函数
func setToken(c *gin.Context, user GINconfig.UserInfo) {
	j := MiddleWare.NewJWT()
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := MiddleWare.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			//ExpiresAt: time.Now().Unix() + 7200,
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Yamada",
		},
	}
	tokenString, err := j.CreatToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "生成u-token出错",
			"token":   nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "token生成成功",
			"data":    user.Username,
			"token":   tokenString,
		})
	}
}
