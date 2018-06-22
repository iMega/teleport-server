package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/imega/teleport-server/config"
	"github.com/imega/teleport-server/health"
	"github.com/imega/teleport-server/shutdown"
	"github.com/imega/teleport-server/uuid"
	"github.com/sirupsen/logrus"
)

// Datastore is a interface store
type Datastore interface {
	CreateEntity(context.Context, Entity) (Entity, error)
	GetEntityByID(context.Context, string) (Entity, error)
	RemoveEntity(context.Context, uuid.UID) error

	CreateRelation(ctx context.Context, subject uuid.UID, predicate string, object uuid.UID, priority int) error
	DeleteRelation(ctx context.Context, subject, object uuid.UID) error

	HealthCheckFunc() health.HealthCheckFunc
	ShutdownFunc() shutdown.ShutdownFunc
}

var (
	marshaller = jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	unmarshaller = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
)

type Entity interface {
	Reset()
	String() string
	ProtoMessage()
	GetId() string
}

type entityDB struct {
	conn   *sql.DB
	logger *logrus.Entry

	create  *sql.Stmt
	update  *sql.Stmt
	remove  *sql.Stmt
	getByID *sql.Stmt

	createRelation *sql.Stmt
	deleteRelation *sql.Stmt
}

func (db *entityDB) HealthCheckFunc() health.HealthCheckFunc {
	return func() bool {
		if err := db.conn.Ping(); err != nil {
			db.logger.Errorf("health: failed to ping database, %s", err)
			return false
		}

		if db.create == nil {
			if err := db.setParepares(); err != nil {
				db.logger.Error(err)
				return false
			}
		}
		return true
	}
}

func (db *entityDB) ShutdownFunc() shutdown.ShutdownFunc {
	return func() {
		if err := db.conn.Close(); err != nil {
			db.logger.Errorf("failed to close connection mysql-server, %s", err)
		}

		db.logger.Info("connection to mysql-server is closed")
	}
}

func NewEntityDB(l *logrus.Entry) (Datastore, error) {
	user, err := config.GetConfigValue("MYSQL_USER")
	if err != nil {
		return nil, fmt.Errorf("failed getting env user, %s", err)
	}

	pass, err := config.GetConfigValue("MYSQL_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("failed getting env password, %s", err)
	}

	host, err := config.GetConfigValue("MYSQL_HOST")
	if err != nil {
		return nil, fmt.Errorf("failed getting env host, %s", err)
	}

	port, err := config.GetConfigValue("MYSQL_PORT")
	if err != nil {
		return nil, fmt.Errorf("failed getting env port, %s", err)
	}

	dbname, err := config.GetConfigValue("MYSQL_DBNAME")
	if err != nil {
		return nil, fmt.Errorf("failed getting env dbname, %s", err)
	}

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pass, host, port, dbname))
	if err != nil {
		return nil, fmt.Errorf("failed to set driver, %s", err)
	}

	db := &entityDB{
		conn:   conn,
		logger: l,
	}

	db.HealthCheckFunc()()

	return db, nil
}

func (db *entityDB) setParepares() error {
	var err error

	if db.create, err = db.conn.Prepare("INSERT INTO entities(owner_id,entity_id,entity_type,entity)VALUES(?,?,?,?)"); err != nil {
		return fmt.Errorf("failed to prepare insert query, %s", err)
	}

	if db.update, err = db.conn.Prepare("UPDATE entities SET entity=? WHERE owner_id=? AND entity_id=? AND removed=0"); err != nil {
		return fmt.Errorf("failed to prepare update query, %s", err)
	}

	if db.remove, err = db.conn.Prepare("UPDATE entities SET removed=1 WHERE owner_id=? AND entity_id=? AND removed=0"); err != nil {
		return fmt.Errorf("failed to prepare remove query, %s", err)
	}

	if db.getByID, err = db.conn.Prepare("SELECT entity_type, entity FROM entities WHERE entity_id = ? AND removed=0"); err != nil {
		return fmt.Errorf("failed to prepare getByID query, %s", err)
	}

	if db.createRelation, err = db.conn.Prepare("INSERT INTO relations(owner_id,subject_id,predicate,object_id,priority)VALUES(?,?,?,?,?)"); err != nil {
		return fmt.Errorf("failed to prepare insert relations query, %s", err)
	}

	if db.deleteRelation, err = db.conn.Prepare("UPDATE relations SET removed=1 WHERE owner_id=? AND subject_id=? AND object_id=? AND removed=0"); err != nil {
		return fmt.Errorf("failed to prepare remove relations query, %s", err)
	}

	return nil
}

func (db *entityDB) CreateEntity(ctx context.Context, e Entity) (Entity, error) {
	ownerID, err := getOwnerIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	entityType := proto.MessageName(e)

	if _, err := marshaller.MarshalToString(e); err != nil {
		return nil, fmt.Errorf("could not create entity, %s", err)
	}
	if _, err := execAffectingOneRow(ctx, db.create, ownerID, uuid.UID(e.GetId()), entityType, "{}"); err != nil {
		return nil, fmt.Errorf("failed to insert entity, %s", err)
	}

	return e, nil
}

func (db *entityDB) GetEntityByID(ctx context.Context, ID string) (Entity, error) {
	var (
		entityType string
		data       []byte
	)

	if err := db.getByID.QueryRowContext(ctx, uuid.UID(ID)).Scan(&entityType, &data); err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, fmt.Errorf("entity not found")
		default:
			return nil, fmt.Errorf("failed getting entity, %s", err)
		}
	}

	r, msg, err := decode(data, entityType)
	if err != nil {
		return nil, fmt.Errorf("failed to decode entity, %s", err)
	}

	if err := unmarshaller.Unmarshal(r, msg); err != nil {
		return nil, fmt.Errorf("failed getting entity, %s", err)
	}

	ent, ok := msg.(Entity)
	if !ok {
		return nil, fmt.Errorf("failed to convert entity, %s", err)
	}

	return ent, nil
}

func (db *entityDB) RemoveEntity(ctx context.Context, entityID uuid.UID) error {
	ownerID, err := getOwnerIDFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err := execAffectingOneRow(ctx, db.remove, ownerID, entityID); err != nil {
		return fmt.Errorf("failed to remove entity, %s", err)
	}

	return nil
}

func execAffectingOneRow(ctx context.Context, stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}

func getOwnerIDFromContext(ctx context.Context) (uuid.UID, error) {
	if id, ok := ctx.Value("owner_id").(string); ok {
		return uuid.UID(id), nil
	}
	return "", fmt.Errorf("failed to extract owner id from context")
}
