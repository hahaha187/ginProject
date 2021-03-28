package controller

import (
  "ginProject/model"
  "ginProject/response"
  "ginProject/vo"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
	"ginProject/common"
)


type IPostController interface {
  RestController
}

type PostController struct {
  DB *gorm.DB
}

func (p PostController) Create(c *gin.Context) {
  var requestPost vo.CreatePostRequst
  if err := c.ShouldBind(&requestPost);err !=nil {
    response.Fail(c,"数据验证错误，分类名称必填",nil)
    return
  }
  //获取登陆用户user
  user,_ :=c.Get("user") // 注册了中间件authmiddle可以从上下文中得到user

  //创建post
  post := model.Post{
    UserId:     user.(model.User).ID, //转换为model。user之后的id
    CategoryId: requestPost.CategoryId,
    Title:      requestPost.Title,
    HeadImg:    requestPost.HeadImg,
    Content:    requestPost.Content,
  }

  if err := p.DB.Create(&post).Error;err != nil {
    panic(err)
    return
  }
  response.Success(c,nil,"创建成功")
}

func (p PostController) Update(c *gin.Context) {
  var requestPost vo.CreatePostRequst
  if err := c.ShouldBind(&requestPost);err !=nil {
    response.Fail(c,"数据验证错误，分类名称必填",nil)
    return
  }

  //获取path中的id
  postId := c.Params.ByName("id")
  var post model.Post
  if p.DB.Where("id = ?",postId).First(&post).RecordNotFound() {
    response.Fail(c,"文章不存在",nil)
    return
  }

  //判断当前用户是否为文章的作者
  //获取登陆用户user
  user,_ :=c.Get("user")
  userId := user.(model.User).ID
  if userId != post.UserId {
    response.Fail(c,"文章不属于您，请勿非法操作",nil)
    return
  }

  //更新文章
  if err := p.DB.Model(&post).Update(requestPost).Error;err != nil {
    response.Fail(c,"更新失败",nil)
    return
  }
  response.Success(c,gin.H{"post":post},"更新成功")


}

func (p PostController) Show(c *gin.Context) {
  //获取path中的id
  postId := c.Params.ByName("id")
  var post model.Post
  //preload gorm wendangkan 19fenzhong
  if p.DB.Preload("Category").Where("id = ?",postId).First(&post).RecordNotFound() {
    response.Fail(c,"文章不存在",nil)
    return
  }
  response.Success(c,gin.H{"post":post},"成功")
}

func (p PostController) Delete(c *gin.Context) {
  //获取path中的id
  postId := c.Params.ByName("id")
  var post model.Post
  if p.DB.Where("id = ?",postId).First(&post).RecordNotFound() {
    response.Fail(c,"文章不存在",nil)
    return
  }

  //判断当前用户是否为文章的作者
  //获取登陆用户user
  user,_ :=c.Get("user")
  userId := user.(model.User).ID
  if userId != post.UserId {
    response.Fail(c,"文章不属于您，请勿非法操作",nil)
    return
  }
  p.DB.Delete(&post)
  response.Success(c,gin.H{"post":post},"删除成功")
}

func NewPostController() IPostController {
  db := common.GetDB()
  db.AutoMigrate(model.Post{})
  return PostController{DB: db}
}
