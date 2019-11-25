package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)
var AppSecret = "newjwt"
type User struct {
	Username string
	Password string
}

type userStdClaims struct {
	jwt.StandardClaims
	*User
}

//实现`type claims interface` 的 `Valid() error`方法  完成自定义检验内容  他会覆盖默认写的Valid-->claims.go
func (c userStdClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(),true) == false{
		return errors.New("token is expried")
	}
	return
}

func JwtGenerateToken(m *User,d time.Duration) (string, error)  {
	expireTime := time.Now().Add(d)  //持续时间
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "zhoujielun",
	}
	ustdclaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,ustdclaims)
	tokenstring, err := token.SignedString([]byte(AppSecret))
	if err != nil{
		panic(err)
	}
	return tokenstring, err
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(AppSecret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
