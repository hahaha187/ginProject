package common

import (
  "ginProject/model"
  "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB{
  dsn := "root:lhy871601318@tcp/gin?charset=utf8&parseTime=True&loc=Local"
  db,err := gorm.Open("mysql",dsn)
  if err!= nil {
    panic("failed to connet database,err:" + err.Error())
  }
  db.AutoMigrate(&model.User{}) // 创建表
  DB = db
  return db
}

func GetDB() *gorm.DB {
  return DB
}
