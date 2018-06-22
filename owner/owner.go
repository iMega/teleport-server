package owner

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/uuid"
)

func GetOwnerIDFromContext(ctx context.Context) (uuid.UID, error) {
	if id, ok := ctx.Value("owner_id").(string); ok {
		return uuid.UID(id), nil
	}
	return "", fmt.Errorf("failed to extract owner id from context")
}
