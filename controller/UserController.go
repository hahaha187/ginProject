package controller

import (
  "ginProject/common"
  "ginProject/dto"
  "ginProject/model"
  "ginProject/response"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
  "log"
  "math/rand"
  "net/http"
  "time"
)

func Register (c *gin.Context) {
  DB := common.GetDB()
  //获取参数
  name := c.PostForm("name")
  telephone := c.PostForm("telephone")
  password := c.PostForm("password")
  //数据验证
  if len(telephone) != 11 {
    response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
    return
  }
  if len(password) < 6 {
    response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
    return
  }
  //如果名称没有传，给一个10位随机的字符串
  if len(name) == 0 {
    name =RandomString(10)
  }
  log.Println(name,telephone,password)
  //判断手机号是否存在
  if isTelephoneExist(DB,telephone) {
    c.JSON(http.StatusUnprocessableEntity,gin.H{
      "code":422,
      "msg":"用户已存在",
    })
    return
  }
  //创建用户
  hasedPassword, err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)  //以给定的Cost返回密码的bcrypt哈希。如果给定的成本小于MinCost，则将成本设置为DefaultCost（10）
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity,gin.H{
      "code":422,
      "msg":"加密错误",
    })
    return
  }
  newUser := model.User{
    Name:name,
    Telephone: telephone,
    Password: string(hasedPassword),
  }
  DB.Create(&newUser)

  //返回结果
  c.JSON(http.StatusOK,gin.H{
    "code":422,
    "message":"注册成功",
  })
}


func Login(c *gin.Context) {
  DB := common.GetDB()
  //获取参数
  telephone := c.PostForm("telephone")
  password := c.PostForm("password")
  //数据验证
  if len(telephone) != 11 {
    c.JSON(http.StatusUnprocessableEntity,gin.H{
      "code":422,
      "msg":"手机号必须为11位",
    })
  }
  if len(password) < 6 {
    c.JSON(http.StatusUnprocessableEntity,gin.H{
      "code":422,
      "msg":"密码不能少于6位",
    })
  }
  //判断手机号是否存在
  var user model.User
  DB.Where("telephone = ?",telephone).First(&user)
  if user.ID == 0 {
    c.JSON(http.StatusUnprocessableEntity,gin.H{
      "code":422,
      "msg":"用户不存在",
    })
  }
  //判断密码是否正确
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)) ; err != nil {//CompareHashAndPassword()用于比对bcrypt哈希字符串和提供的密码明文文本是否匹配。
    c.JSON(http.StatusBadRequest,gin.H{
      "coed":400,
      "msg": "密码错误",
    })
    return
  }
  //发放token
  token,err := common.ReleaseToken(user)
  if err != nil {
    c.JSON(http.StatusInternalServerError,gin.H{
      "code":500,
      "msg": "系统异常",
    })
    log.Printf("token generate error :%v",err)
    return
  }
  //返回结果
  c.JSON(http.StatusUnprocessableEntity,gin.H{
    "code":200,
    "data": gin.H{
      "token" :token,
    },
    "msg":"登陆成功",
  })

}


func Info(c *gin.Context) {
  user,_ := c.Get("user")
  c.JSON(http.StatusOK,gin.H{
    "code":200,
    "data":gin.H{"user":dto.ToUserDto(user.(model.User))},//这里的user是
  })
}

func isTelephoneExist(db *gorm.DB,telephone string) bool{
  var user model.User
  db.Where("telephone = ?",telephone).First(&user)
  if user.ID != 0 {
    return true
  }
  return false
}

func RandomString(n int) string {
  var letters = []byte("asdfghjklqwerttyuzxcvbnm")
  result := make([]byte,n)
  rand.Seed(time.Now().Unix()) // 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
  for i:= range result {
    result[i] = letters[rand.Intn(len(letters))]
  }
  return string(result)
}
