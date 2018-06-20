package api

import "github.com/golang/protobuf/proto"

func init() {
	proto.RegisterType((*Token)(nil), "ru.imega.teleport.token")
	proto.RegisterType((*Product)(nil), "ru.imega.teleport.product")
}
