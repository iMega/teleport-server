package resolver

import (
	"testing"
)

func Test_FromGetter_GetString_Ptr_ReturnsValue(t *testing.T) {
	type Input struct {
		Pass *string
	}

	expect := "qweqwe111"
	in := Input{
		Pass: &expect,
	}

	actual, ok := FromGetter(in.Pass).GetString()
	if !ok {
		t.Error("failed getting value from GetString()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetString_ReturnsValue(t *testing.T) {
	type Input struct {
		Pass string
	}

	expect := "qweqwe111"
	in := Input{
		Pass: expect,
	}

	actual, ok := FromGetter(in.Pass).GetString()
	if !ok {
		t.Error("failed getting value from GetString()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetString_PtrEmptyValue(t *testing.T) {
	type Input struct {
		Pass *string
	}

	in := Input{}
	_, ok := FromGetter(in.Pass).GetString()
	if ok {
		t.Error("value must be empty")
	}
}

func Test_FromGetter_GetString_EmptyValue(t *testing.T) {
	type Input struct {
		Pass string
	}

	in := Input{}
	_, ok := FromGetter(in.Pass).GetString()
	if ok {
		t.Error("value must be empty")
	}
}

func Test_FromGetter_GetInt32_Ptr_ReturnsValue(t *testing.T) {
	type Input struct {
		Integer *int32
	}

	expect := int32(100)
	in := Input{
		Integer: &expect,
	}

	actual, ok := FromGetter(in.Integer).GetInt32()
	if !ok {
		t.Error("failed getting value from GetInt32()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetInt32_ReturnsValue(t *testing.T) {
	type Input struct {
		Integer int32
	}

	expect := int32(100)
	in := Input{
		Integer: expect,
	}

	actual, ok := FromGetter(in.Integer).GetInt32()
	if !ok {
		t.Error("failed getting value from GetInt32()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetInt32_PtrEmptyValue(t *testing.T) {
	type Input struct {
		Integer *int32
	}

	in := Input{}
	_, ok := FromGetter(in.Integer).GetInt32()
	if ok {
		t.Error("value must be empty")
	}
}

func Test_FromGetter_GetInt32_EmptyValue(t *testing.T) {
	type Input struct {
		Integer int32
	}

	in := Input{}
	_, ok := FromGetter(in.Integer).GetInt32()
	if ok {
		t.Error("value must be empty")
	}
}

func Test_FromGetter_GetInt64_Ptr_ReturnsValue(t *testing.T) {
	type Input struct {
		Integer *int64
	}

	expect := int64(100)
	in := Input{
		Integer: &expect,
	}

	actual, ok := FromGetter(in.Integer).GetInt64()
	if !ok {
		t.Error("failed getting value from GetInt64()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetInt64_ReturnsValue(t *testing.T) {
	type Input struct {
		Integer int64
	}

	expect := int64(100)
	in := Input{
		Integer: expect,
	}

	actual, ok := FromGetter(in.Integer).GetInt64()
	if !ok {
		t.Error("failed getting value from GetInt64()")
	}

	if *actual != expect {
		t.Error("values is not equals")
	}
}

func Test_FromGetter_GetInt64_PtrEmptyValue(t *testing.T) {
	type Input struct {
		Integer *int64
	}

	in := Input{}
	_, ok := FromGetter(in.Integer).GetInt64()
	if ok {
		t.Error("value must be empty")
	}
}

func Test_FromGetter_GetInt64_EmptyValue(t *testing.T) {
	type Input struct {
		Integer int64
	}

	in := Input{}
	_, ok := FromGetter(in.Integer).GetInt64()
	if ok {
		t.Error("value must be empty")
	}
}
