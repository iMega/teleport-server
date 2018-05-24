package mysql

import (
	"context"
	"fmt"

	"github.com/imega/teleport-server/uuid"
)

func (db *entityDB) CreateRelation(ctx context.Context, subject uuid.UID, predicate string, object uuid.UID, priority int) error {
	ownerID, err := getOwnerIDFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err := execAffectingOneRow(ctx, db.createRelation, ownerID, subject, predicate, object, priority); err != nil {
		return fmt.Errorf("failed to create relation, %s", err)
	}

	return nil
}

func (db *entityDB) DeleteRelation(ctx context.Context, subject, object uuid.UID) error {
	return nil
}
