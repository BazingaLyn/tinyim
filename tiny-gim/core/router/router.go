package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SendReq struct {
	Id      string `form:"id"`
	Content string `form:"content"`
}

type HttpResp struct {
	Success bool
	Msg     string
}

func Connect(c *gin.Context) {

}

func Send(c *gin.Context) {

	var sendReq SendReq

	if c.ShouldBind(&sendReq) == nil {
		fmt.Println(sendReq)
	}

	c.String(200, "ok")

	//receiverId := c.Query("id")
	//fmt.Println(c.Query("content"))
	//msg := []byte(c.Query("content"))
	//
	//if i, err := strconv.ParseInt(receiverId, 10, 64); err == nil {
	//
	//	ok, err := transport.HandlerSendMsg(i, msg)
	//	if err != nil || !ok {
	//		fmt.Println("ok is ",ok ," err is ",err)
	//		c.String(http.StatusOK, "failed")
	//	}
	//} else {
	//	c.String(http.StatusOK, "id is not number")
	//}
	//
	//c.String(http.StatusOK, "ok")

}
