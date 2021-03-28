package load

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoader_Load(t *testing.T) {
	loader, err := NewLoader("../../testdata/type-field.xlsx")
	assert.NoError(t, err)
	assert.NoError(t, loader.Load("DifferentTypes"))

	buf := &bytes.Buffer{}
	assert.NoError(t, loader.OutputProto(buf, "example.foo"))
	assert.EqualValues(t, `syntax = "proto3";

package example.foo;

message DifferentTypes {
  string foo_string = 1;
  uint32 foo_uint32 = 2;
  int32 foo_int32 = 3;
  uint64 foo_uint64 = 4;
  int64 foo_int64 = 5;
  repeated string foo_list_string = 6;
  repeated int64 foo_list_int64 = 7;
  BarStruct foo_list_struct = 8;
}

message foo_list_struct {
  int64 id = 1;
  int64 cnt = 2;
  int64 id = 1;
  int64 cnt = 2;
}
`, buf.String())
}
