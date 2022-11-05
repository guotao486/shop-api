/*
 * @Author: GG
 * @Date: 2022-11-05 15:58:09
 * @LastEditTime: 2022-11-05 17:03:29
 * @LastEditors: GG
 * @Description: main
 * @FilePath: \shop-api\main.go
 *
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping/api"
	_ "shopping/docs"
	"shopping/utils/graceful"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSagger "github.com/swaggo/gin-swagger"
)

// @title 电商项目
// @description 电商项目
// @version 1.0
// @contact.name golang技术栈
// @contact.url https://www.golang-tech-stack.com

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	registerMiddlewares(r)
	api.RegisterHandlers(r)
	r.GET("/swagger/*any", ginSagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	graceful.ShutdownGin(srv, time.Second*5)
}

func registerMiddlewares(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(params gin.LogFormatterParams) string {
				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s %d %s %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC3339),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.ErrorMessage,
				)
			}))

	r.Use(gin.Recovery())
}
