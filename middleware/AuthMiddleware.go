package middleware

import (
  "ginProject/common"
  "ginProject/model"
  "github.com/gin-gonic/gin"
  "net/http"
  "strings"
)

func AuthMiddleware () gin.HandlerFunc { //gin 的中间件就是一个函数返回gin.HandlerFunc
  return func(c *gin.Context) {
    //获取authorization header
    tokenString :=c.GetHeader("Authorization")

    //验证格式
    if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer ") {
      c.JSON(http.StatusUnauthorized,gin.H{
        "code":401,
        "msg":"权限不足",
      })
      c.Abort()  //阻止后续执行
      return
    }

    tokenString = tokenString[7:]

    token,claims,err := common.ParseToken(tokenString)
    if err != nil || !token.Valid{
      c.JSON(http.StatusUnauthorized,gin.H{
        "code":401,
        "msg":"权限不足",
      })
      c.Abort()
      return
    }

    //验证通过后获取claim中的userId
    userId := claims.UserId
    DB := common.GetDB()
    var user model.User
    DB.First(&user,userId)

    //yonghu
    if user.ID == 0 {
      c.JSON(http.StatusUnauthorized,gin.H{
        "code":401,
        "msg":"权限不足",
      })
      c.Abort()
      return
    }

    //用户存在将user的信息写入上下文
    c.Set("user",user)
    c.Next()
  }
}
