package node

import (
	"github.com/golang/protobuf/proto"
	"github.com/imega/teleport-server/mysql"
)

// Resolver is a interface node of resolver
type Resolver interface {
	setDatastore(mysql.Datastore)
	setResolvable(item interface{})
}

// Option is a option of resolver
type Option func(r Resolver)

// WithDatastore set datastore
func WithDatastore(ds mysql.Datastore) Option {
	return func(r Resolver) {
		r.setDatastore(ds)
	}
}

// WithNode ...
func WithNode(node proto.Message) Option {
	return func(r Resolver) {
		r.setResolvable(node)
	}
}
