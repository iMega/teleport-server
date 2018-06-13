package node

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/mysql"
)

// UserResolver ...
type UserResolver struct {
	user *api.User
	ds   mysql.Datastore
}

func (r *UserResolver) setDatastore(ds mysql.Datastore) {
	r.ds = ds
}
func (r *UserResolver) setResolvable(item interface{}) {
	r.user = item.(*api.User)
}

// NewUserResolver ...
func NewUserResolver(opts ...Option) *UserResolver {
	r := &UserResolver{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// ID возвращает значение поля id
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.GetId())
}

// Active возвращает значение поля active
func (r *UserResolver) Active() bool {
	return r.user.GetActive()
}

// Pass возвращает значение поля active
func (r *UserResolver) Pass() string {
	return "*****"
}
