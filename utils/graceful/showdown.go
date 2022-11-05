/*
 * @Author: GG
 * @Date: 2022-11-05 15:43:01
 * @LastEditTime: 2022-11-05 15:54:52
 * @LastEditors: GG
 * @Description: 优雅的关闭gin服务器工具类
 * @FilePath: \shop-api\utils\graceful\showdown.go
 *
 */
package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
 * @description: 优雅的关闭gin服务工具类
 * @param {*http.Server} instance
 * @param {time.Duration} timeout
 * @return {*}
 */
func ShowdownGin(instance *http.Server, timeout time.Duration) {
	// 操作信号
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 接收通道
	<-quit

	log.Println("关闭 service......")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("Service 关闭:", err)
	}

	// 超时5秒，ctx.Done()
	select {
	case <-ctx.Done():
		log.Println("超时5秒")
	}
	log.Println("Service 退出")
}
