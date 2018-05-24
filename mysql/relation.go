package mysql

import (
	"context"

	"github.com/imega/teleport-server/uuid"
)

func (db *entityDB) CreateRelation(ctx context.Context, subject uuid.UID, predicate string, object uuid.UID, priority int) error {
	return nil
}

func (db *entityDB) DeleteRelation(ctx context.Context, subject, object uuid.UID) error {
	return nil
}
