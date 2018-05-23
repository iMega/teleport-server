package api

import "github.com/golang/protobuf/proto"

func init() {
	proto.RegisterType((*Product)(nil), "ru.imega.teleport.product")
}
