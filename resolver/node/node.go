package node

import (
	"github.com/golang/protobuf/proto"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/mysql"
)

// NodeResolver резолвер ноды
type NodeResolver struct {
	datastore mysql.Datastore
	node      proto.Message
}

// NewNodeResolver getting new instance resolver
func NewNodeResolver(opts ...Option) *NodeResolver {
	r := &NodeResolver{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *NodeResolver) setDatastore(ds mysql.Datastore) {
	r.datastore = ds
}

func (r *NodeResolver) setResolvable(n interface{}) {
	r.node = n.(proto.Message)
}

func (r *NodeResolver) ID() graphql.ID {
	if e, ok := r.node.(mysql.Entity); ok {
		return graphql.ID(e.GetId())
	}
	return graphql.ID("")
}

func (r *NodeResolver) ToProduct() (*ProductResolver, bool) {
	if p, ok := r.node.(*api.Product); ok {
		return NewProductResolver(
			WithDatastore(r.datastore),
			WithNode(p),
		), true
	}

	return nil, false
}

func (r *NodeResolver) ToUser() (*UserResolver, bool) {
	if p, ok := r.node.(*api.User); ok {
		return NewUserResolver(
			WithDatastore(r.datastore),
			WithNode(p),
		), true
	}

	return nil, false
}
