package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/assert"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resource"
	"go.uber.org/zap"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.5.9
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		err := global.GVA_DB.AutoMigrate(example.ExaFile{},
			example.ExaCustomer{},
			example.ExaFileChunk{},
			example.ExaFileUploadAndDownload{}, resource.Video{}, resource.Image{}, assert.Target{}, ad.Plan{}, assert.Policy{}, resource.Material{}, ad.Campaign{}, assert.BlackWhiteList{}, ad.Creative{}) // 初始化表
		if err != nil {
			global.GVA_LOG.Error("register table failed", zap.Error(err))
			os.Exit(0)
		}
		global.GVA_LOG.Info("register table success")

		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()

		bid.Load()
		initialize.BidTimer()
	}
	dsp.RunServer()
}
