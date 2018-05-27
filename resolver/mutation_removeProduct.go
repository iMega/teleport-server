package resolver

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/uuid"
)

// RemoveProductInput входные аргументы
type RemoveProductInput struct {
	ID uuid.UID
}

// RemoveProduct убрать товар
func (r *Resolver) RemoveProduct(ctx context.Context, args RemoveProductInput) (bool, error) {
	if err := r.EntityDB.RemoveEntity(ctx, args.ID); err != nil {
		return false, fmt.Errorf("failed to remove entity, %s", err)
	}
	return true, nil
}
