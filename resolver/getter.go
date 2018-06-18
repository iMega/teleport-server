package resolver

import (
	"reflect"
)

// Getter проверка существования и получение значений
type Getter interface {
	GetString() (*string, bool)
	GetInt32() (*int32, bool)
	GetInt64() (*int64, bool)
}

// FromGetter ...
func FromGetter(i interface{}) Getter {
	return &getter{
		val: i,
	}
}

type getter struct {
	val interface{}
}

func (g *getter) isZero() bool {
	return g.val == reflect.Zero(reflect.TypeOf(g.val)).Interface()
}

func (g *getter) GetString() (*string, bool) {
	if g.isZero() {
		return nil, false
	}
	val := reflect.Indirect(reflect.ValueOf(g.val))
	if a, ok := val.Interface().(string); ok {
		return &a, true
	}
	return nil, false
}

func (g *getter) GetInt32() (*int32, bool) {
	if g.isZero() {
		return nil, false
	}
	val := reflect.Indirect(reflect.ValueOf(g.val))
	if a, ok := val.Interface().(int32); ok {
		return &a, true
	}
	return nil, false
}

func (g *getter) GetInt64() (*int64, bool) {
	if g.isZero() {
		return nil, false
	}
	val := reflect.Indirect(reflect.ValueOf(g.val))
	if a, ok := val.Interface().(int64); ok {
		return &a, true
	}
	return nil, false
}
