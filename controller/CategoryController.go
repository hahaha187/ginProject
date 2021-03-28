package controller

import (
 // "ginProject/common"
  "ginProject/model"
  "ginProject/repository"

  //repository2 "ginProject/repository"
  "ginProject/response"
  "ginProject/vo"
  "github.com/gin-gonic/gin"
 // "github.com/jinzhu/gorm"
  "strconv"
)

type ICategoryController interface { //定义一个接口
  RestController
}

type CategoryController struct {
  Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController{
  repository := repository.NewCategoryRepository()
  repository.DB.AutoMigrate(model.Category{})
  return CategoryController{Repository: repository}
}

func (c2 CategoryController) Create(c *gin.Context) {

  var requestCategory vo.CreateCategoryRequest
  if err := c.ShouldBind(&requestCategory);err !=nil {
    response.Fail(c,"数据验证错误，分类名称必填",nil)
    return
  }
  //category := model.Category{Name: requestCategory.Name}
  //c2.DB.Create(&category) //创建分类
  category,err :=c2.Repository.Create(requestCategory.Name)
  if  err !=nil {
    //response.Fail(c,"创建失败",nil)
    panic(err)
    return
  }

  //response.Success(c,gin.H{"category":requestCategory},"")
  response.Success(c,gin.H{"category":category},""  )
}

func (c2 CategoryController) Update(c *gin.Context) {
  //绑定body中的参数
  //var requestCategory vo.CreateCategoryRequest
  //if err := c.ShouldBind(&requestCategory);err !=nil {
  //  response.Fail(c,"数据验证错误，分类名称必填",nil)
  //  return
  //}
  var requestCategory vo.CreateCategoryRequest
  if err := c.ShouldBind(&requestCategory);err !=nil {
    response.Fail(c,"数据验证错误，分类名称必填",nil)
    return
  }
  //获取path中的参数
  categoryId,_ := strconv.Atoi(c.Params.ByName("id"))
  //var updateCategory model.Category
  //if c2.DB.First(&updateCategory,categoryId).RecordNotFound() {
  //  response.Fail(c,"分类不存在",nil)
  //}

  updateCategory,err := c2.Repository.SelectById(categoryId)
  if err != nil {
    response.Fail(c,"分类不存在",nil)
    return
  }

  //更新分类
  category,err := c2.Repository.Update(*updateCategory,requestCategory.Name)
  if err != nil {
    panic(err)
  }
  //c2.DB.Model(&updateCategory).Update("name",requestCategory.Name)
  //response.Success(c,gin.H{"category":updateCategory},"修改成功")
  response.Success(c,gin.H{"category":category},"修改成功")

}

func (c2 CategoryController) Show(c *gin.Context) {
  //获取path中的参数
  categoryId,_ := strconv.Atoi(c.Params.ByName("id"))
  //var category model.Category
  //if c2.DB.First(&category,categoryId).RecordNotFound() {
  //  response.Fail(c,"分类不存在",nil)
  //  return
  //}
  category,err := c2.Repository.SelectById(categoryId)
  if err != nil {
    response.Fail(c,"分类不存在",nil)
    return
  }
  response.Success(c,gin.H{"category":category},"")
}

func (c2 CategoryController) Delete(c *gin.Context) {
  //获取path中的参数
  categoryId,_ := strconv.Atoi(c.Params.ByName("id"))
  //var category model.Category
  //if c2.DB.First(&category,categoryId).RecordNotFound() {
  //  response.Fail(c,"分类不存在",nil)
  //}
  //if err := c2.DB.Delete(model.Category{},categoryId).Error;err !=nil {
  // response.Fail(c,"删除失败，请重试",nil)
  // return
  //}

  if err := c2.Repository.DeleteById(categoryId);err !=nil {
    response.Fail(c,"删除失败，请重试",nil)
    return
  }

  response.Success(c,nil,"")
}

