package controller

import "github.com/gin-gonic/gin"

type RestController interface {  //定义一个接口（只定义规范不实现，由具体的对象来实现规范的细节）接口是一种类型
  Create(c *gin.Context)
  Update(c *gin.Context)
  Show(c *gin.Context)
  Delete(c *gin.Context)
}
