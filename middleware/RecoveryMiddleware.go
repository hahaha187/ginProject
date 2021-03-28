package middleware

import (
  "fmt"
  "ginProject/response"
  "github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    defer func() { // defer后面是匿名函数时，要在结尾加(),表示调用该匿名函数，
      if err := recover(); err != nil {
        response.Fail(c, fmt.Sprint(err), nil)
      }
    }()
    c.Next()
  }
}
