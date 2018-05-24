package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/uuid"
)

type DeleteProductInput struct {
	ID uuid.UID
}

// DeleteProduct удаление товара
func (r *Resolver) DeleteProduct(ctx context.Context, args DeleteProductInput) (bool, error) {
	if err := r.EntityDB.DeleteEntity(ctx, args.ID); err != nil {
		return false, fmt.Errorf("failed to delete entity, %s", err)
	}
	return true, nil
}
