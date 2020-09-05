package codec

import (
	"encoding/json"
	"fmt"
	"testing"
	"tiny-gim/transport"
)

func TestMessagePackCodec(t *testing.T) {

	messagePackCodec := new(MessagePackCodec)

	loginReq := &transport.LoginReq{
		Id:   "21",
		Name: "Bazinga",
		Pwd:  "123456",
	}
	bytes, err := messagePackCodec.Encode(loginReq)

	if err != nil {
		fmt.Println("encode is failed ", err)
		return
	}

	loginReq2 := &transport.LoginReq{}

	messagePackCodec.Decode(bytes, loginReq2)

	marshal, err := json.Marshal(loginReq2)

	if err == nil {
		t.Log("print result")
		fmt.Println(string(marshal))
	}

}
