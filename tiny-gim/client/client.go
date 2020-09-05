package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tiny-gim/metadata"
	"tiny-gim/transport/codec"
	"tiny-gim/transport/protocol"
)

var messagePackCodec = new(codec.MessagePackCodec)

var imProtocol = new(protocol.IMProtocol)

func main() {

	header := &protocol.Header{
		Seq:     1,
		MsgType: protocol.MessageLoginRequest,
	}

	loginReq := &metadata.LoginReq{
		Id:   "21",
		Name: "Bazinga",
		Pwd:  "123456",
	}

	bytes, err := messagePackCodec.Encode(loginReq)

	if err != nil {
		fmt.Println("encode is failed ", err)
		return
	}

	message := &protocol.Message{
		Header: header,
		Data:   bytes,
	}

	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {
		panic(err)
	}

	encodeMessage := imProtocol.EncodeMessage(message)
	//var buf = make([]byte, 8)
	//	//binary.BigEndian.PutUint64(buf, uint64(1))

	if _, err := conn.Write(encodeMessage); err != nil {
		fmt.Println(err)
	}

	for {
		buffer := make([]byte, 1024)
		recvNum, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				// client 连接关闭
				break
			}
			fmt.Println("count is {}", recvNum)
			fmt.Println("err recv from server: ", err)
		} else {
			msg := string(buffer[:recvNum])
			fmt.Println("recv from server: ", msg)
		}

		time.Sleep(time.Second * 2)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
	<-ch
}
