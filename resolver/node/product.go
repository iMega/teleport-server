package node

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/mysql"
)

// ProductResolver ...
type ProductResolver struct {
	product *api.Product
	ds      mysql.Datastore
}

func (r *ProductResolver) setDatastore(ds mysql.Datastore) {
	r.ds = ds
}
func (r *ProductResolver) setResolvable(item interface{}) {
	r.product = item.(*api.Product)
}

// NewProductResolver ...
func NewProductResolver(opts ...Option) *ProductResolver {
	r := &ProductResolver{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *ProductResolver) ID() graphql.ID {
	return graphql.ID(r.product.GetId())
}

func (r *ProductResolver) Title() *string {
	ret := r.product.GetTitle()
	return &ret
}
