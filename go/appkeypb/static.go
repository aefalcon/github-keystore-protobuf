package appkeypb

import (
	"github.com/golang/protobuf/proto"
)

var DefaultLinks Links
var DefaultLinksPb = []byte{10, 9, 105, 110, 100, 101, 120, 46, 112, 98, 102, 18, 15, 97, 112, 112, 47, 123, 65, 112, 112, 73, 100, 125, 46, 112, 98, 102, 26, 29, 107, 101, 121, 47, 123, 65, 112, 112, 73, 100, 125, 47, 123, 70, 105, 110, 103, 101, 114, 112, 114, 105, 110, 116, 125, 46, 98, 105, 110, 34, 34, 107, 101, 121, 47, 123, 65, 112, 112, 73, 100, 125, 47, 123, 70, 105, 110, 103, 101, 114, 112, 114, 105, 110, 116, 125, 45, 109, 101, 116, 97, 46, 112, 98, 102}

func init() {
	err := proto.Unmarshal(DefaultLinksPb, &DefaultLinks)
	if err != nil {
		panic(err.Error())
	}
}
