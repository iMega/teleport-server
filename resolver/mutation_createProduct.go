package resolver

import (
	"context"

	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/resolver/node"
	"github.com/imega/teleport-server/uuid"
)

// CreateProduct ...
func (r *Resolver) CreateProduct(ctx context.Context) (*node.ProductResolver, error) {
	entity, err := r.EntityDB.CreateEntity(ctx, &api.Product{
		Id: string(uuid.NewUUID()),
	})
	if err != nil {
		return nil, err
	}

	return node.NewProductResolver(node.WithNode(entity), node.WithDatastore(r.EntityDB)), nil
}
