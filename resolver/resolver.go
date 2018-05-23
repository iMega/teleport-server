package resolver

import "github.com/imega/teleport-server/mysql"

type Resolver struct {
	EntityDB mysql.Datastore
}
