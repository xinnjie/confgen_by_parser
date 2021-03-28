package proto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestRestore(t *testing.T) {
	var (
		syntax      = "proto3"
		packageName = "example.foo"
		messageName = "Foo"

		labelOptional = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
		labelRepeated = descriptorpb.FieldDescriptorProto_LABEL_REPEATED

		uint32fieldName         = "foo_uint32"
		uint32fieldNumber int32 = 1
		uint32Type              = descriptorpb.FieldDescriptorProto_TYPE_UINT32

		int32fieldName         = "foo_int32"
		int32fieldNumber int32 = 2
		int32Type              = descriptorpb.FieldDescriptorProto_TYPE_INT32

		uint64fieldName         = "foo_uint64"
		uint64fieldNumber int32 = 3
		uint64Type              = descriptorpb.FieldDescriptorProto_TYPE_UINT64

		int64fieldName         = "foo_int64"
		int64fieldNumber int32 = 4
		int64Type              = descriptorpb.FieldDescriptorProto_TYPE_INT64

		repeatedStringFieldName         = "foo_list_string"
		repeatedStringFieldNumber int32 = 5
		stringType                      = descriptorpb.FieldDescriptorProto_TYPE_STRING

		proto = &descriptorpb.FileDescriptorProto{
			MessageType: []*descriptorpb.DescriptorProto{
				{
					Name: &messageName,
					Field: []*descriptorpb.FieldDescriptorProto{
						{
							Name:   &uint32fieldName,
							Number: &uint32fieldNumber,
							Label:  &labelOptional,
							Type:   &uint32Type,
						},
						{
							Name:   &int32fieldName,
							Number: &int32fieldNumber,
							Label:  &labelOptional,
							Type:   &int32Type,
						},
						{
							Name:   &uint64fieldName,
							Number: &uint64fieldNumber,
							Label:  &labelOptional,
							Type:   &uint64Type,
						},
						{
							Name:   &int64fieldName,
							Number: &int64fieldNumber,
							Label:  &labelOptional,
							Type:   &int64Type,
						},
						{
							Name:   &repeatedStringFieldName,
							Number: &repeatedStringFieldNumber,
							Label:  &labelRepeated,
							Type:   &stringType,
						},
					},
				},
			},
			Package: &packageName,
			Syntax:  &syntax,
		}
	)

	w := &bytes.Buffer{}
	Restore(w, proto)
	assert.EqualValues(t, `syntax = "proto3";

package example.foo;

message Foo {
  uint32 foo_uint32 = 1;
  int32 foo_int32 = 2;
  uint64 foo_uint64 = 3;
  int64 foo_int64 = 4;
  repeated string foo_list_string = 5;
}
`, w.String())
}
