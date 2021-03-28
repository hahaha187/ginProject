package main

import (
  "ginProject/common"
  "github.com/gin-gonic/gin"
  _"github.com/go-sql-driver/mysql"
  "github.com/spf13/viper"
  "os"
)

func main()  {
  InitConfig()
  db := common.InitDB()
  defer db.Close()
  r:=gin.Default()
  r = CollectRouter(r)
  r.Run(":8080")
}
//从文件中读取配置
func InitConfig() {
  workDir,_ :=os.Getwd()
  viper.SetConfigName("application")
  viper.SetConfigType("yml")
  viper.AddConfigPath(workDir + "/config")
  err := viper.ReadInConfig()
  if err!=nil {
    panic(err)
  }
}




