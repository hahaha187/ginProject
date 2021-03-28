package common

import (
  "ginProject/model"
  "github.com/dgrijalva/jwt-go"
  "time"
)
//自定义一个key,本地密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
  UserId uint `json:"userid"`
  jwt.StandardClaims
}
//生成jwt
func ReleaseToken(user model.User)(string,error) {
  expirationTime := time.Now().Add(7*24*time.Hour) // token有效时间
  //创建一个我们自己的声明
  claims := &Claims{
    UserId: user.ID,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expirationTime.Unix(), //过期时间
      IssuedAt: time.Now().Unix(),  //发放时间
      Issuer: "lzl", // 签发人
      Subject: "user token",  //项目
    },
  }

  //使用指定的签名方法创建签名对象
  token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
  //使用指定的key签名并获得完整的编码后的字符串
  tokenString,err := token.SignedString(jwtKey)
  if err != nil {
    return "", err
  }
  return tokenString,nil
}
//解析jwt
func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
  claims := &Claims{}
  //解析token
  token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{},err error) {
    return jwtKey,nil
  })
  return token,claims,err
}

