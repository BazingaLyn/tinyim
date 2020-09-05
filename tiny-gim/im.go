package main

import (
	"time"
	"tiny-gim/transport"
)

func main() {

	go func() {
		transport.InitNetwork()
	}()

	time.Sleep(500 * time.Second)
	//engine := gin.Default()
	//
	//// 发送消息
	//engine.GET("/send", router.Send)
	//// 获取请求连接
	//engine.GET("/connect",router.Connect)
	//
	//engine.Run()

}
