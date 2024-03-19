//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"
	ijwt "webook/internal/web/jwt"
	"webook/ioc"
)

func InitWebServer() *gin.Engine {
	//第三方依赖
	wire.Build(
		ioc.InitRedis, ioc.InitDB,
		ioc.InitLogger,
		//Dao 部分
		dao.NewUserDao,
		//cache 部分
		cache.NewUserCache,
		//repository 部分
		repository.NewUserRepository,
		//service部分
		service.NewUserService,
		//handler部分
		ijwt.NewRedisJWTHandler,
		web.NewUserHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWeb,
	)
	return gin.Default()
}
