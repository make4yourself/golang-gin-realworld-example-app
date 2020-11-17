package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	"github.com/wangzitian0/golang-gin-starter-kit/articles"
	"github.com/wangzitian0/golang-gin-starter-kit/common"
	"github.com/wangzitian0/golang-gin-starter-kit/users"
)

// 如果需要就迁移数据库
// GORM:Migration 文档 https://gorm.io/zh_CN/docs/migration.html
func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
}

func main() {

	// 初始化数据库的连接
	db := common.Init()
	Migrate(db)

	// 在程序结束(该主函数结束)或出现 panic 的时候关闭执行
	// 关闭数据库的连接，减轻对数据库服务的压力
	defer db.Close()

	// 初始化 gin
	r := gin.Default()

	// 使用 cors 中间件
	r.Use(cors.Default())

	// 注册（创建） "/api" 路由组
	v1 := r.Group("/api")

	// 注册（创建）用户注册相关的路由在 "/users" 路由组下
	users.UsersRegister(v1.Group("/users"))

	// 将不需要用户认证的逻辑中间件加入 v1 路由组并注册不需要认证操作的路由
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	// 重新添加需要认证的逻辑中间件到 v1 路由组然后注册需要认证的路由路径
	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	articles.ArticlesRegister(v1.Group("/articles"))

	// 该路由使用 r 的 默认路由组所以没有继承之前有中间件注入的路有组，因此没有中间件逻辑会走
	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// test 1 to 1
	tx1 := db.Begin()
	userA := users.UserModel{
		Username: "AAAAAAAAAAAAAAAA",
		Email:    "aaaa@g.cn",
		Bio:      "hehddeda",
		Image:    nil,
	}
	tx1.Save(&userA)
	tx1.Commit()
	fmt.Println(userA)

	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	r.Run() // listen and serve on 0.0.0.0:8080
}
