package User

import (
	//本地
	"encoding/base64"
	"log"
	"net/http"
	//第三方
	"golang.org/x/crypto/scrypt"
	//目录
	"File-sharing-web-page/GINServer/GINconfig"
	"File-sharing-web-page/Global"
)

// CreatUser 创建用户
func CreatUser(data *GINconfig.UserInfo) (code bool) {
	data.Password = ScryptPW(data.Password)
	err := Global.Global.DB.Create(&data).Error
	if err != nil {
		code = false
		return
	} else {
		code = true
	}
	return
}

// BeforeCreate 密码加密 密码加密方法有：bcrypt，scrypt，加salt hash

func ScryptPW(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{255, 22, 18, 33, 99, 66, 25, 11}
	HashPW, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	//HashPW, err :=bcrypt.GenerateFromPassword([]byte(password),KeyLen)

	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPW)
	return fpw
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users GINconfig.UserInfo
	Global.Global.DB.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		code = http.StatusConflict
		return
	} else {
		code = http.StatusOK
	}
	return
}

// ValidateLogin 后台登陆验证
func ValidateLogin(username string, password string) (userinfo GINconfig.UserInfo, code int) {
	var user GINconfig.UserInfo

	Global.Global.DB.Where("username = ?", username).First(&user)
	userinfo = user

	if user.ID == 0 {
		code = http.StatusConflict
		return
	}

	if ScryptPW(password) != user.Password {
		code = http.StatusUnauthorized
		return
	}

	code = http.StatusOK

	return
}
