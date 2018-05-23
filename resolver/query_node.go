package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/resolver/node"
)

// QueryNodeInput входные данные резолвера
type QueryNodeInput struct {
	ID graphql.ID
}

// GetID получение ID
func (i *QueryNodeInput) GetID() string {
	return string(i.ID)
}

// Node получение node
func (r *Resolver) Node(ctx context.Context, in QueryNodeInput) (*node.NodeResolver, error) {
	n, _ := r.EntityDB.GetEntityByID(ctx, in.GetID())
	return node.NewNodeResolver(node.WithNode(n), node.WithDatastore(r.EntityDB)), nil
}
