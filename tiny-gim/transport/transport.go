package transport

import (
	"errors"
	"fmt"
	"net"
	"tiny-gim/metadata"
	"tiny-gim/transport/codec"
	"tiny-gim/transport/protocol"
)

var pool = make(map[int64]net.Conn)

//func init() {
// fmt.Println("network now init")
// initNetwork()
//}

var messagePackCodec = new(codec.MessagePackCodec)

var imProtocol = new(protocol.IMProtocol)

func InitNetwork() {

	lis, err := net.Listen("tcp", "127.0.0.1:10086")

	fmt.Println("hello listener begin")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	msg, err := imProtocol.DecodeMessage(conn)

	if err != nil {
		fmt.Println("decode message err", err)
		conn.Close()
		return
	}

	fmt.Println("handle conn in..")

	switch msg.MsgType {
	case protocol.MessageLoginRequest:
		loginReq := &metadata.LoginReq{}

		if err := messagePackCodec.Decode(msg.Data, loginReq); err != nil {
			fmt.Println("handleConn failed")
			return
		}

		fmt.Println(loginReq)
	}

	//var buf = make([]byte, 8)
	//_, err := io.ReadFull(conn, buf)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//key := int64(binary.BigEndian.Uint64(buf))
	//
	//pool[key] = conn

}

func HandlerSendMsg(key int64, msg []byte) (bool, error) {

	if conn, ok := pool[key]; ok {
		fmt.Println("ready send content size is ", len(msg))
		_, err := conn.Write(msg)
		if err != nil {
			return false, err
		}
	} else {
		fmt.Printf("key %d not find match conn", key)
		return false, errors.New("not find match person")
	}

	return true, nil
}
