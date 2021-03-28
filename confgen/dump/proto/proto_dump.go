package proto

import (
	"fmt"
	"github.com/xinnjie/confgen_by_parser/confgen/ast"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	LabelOptional = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	LabelRepeated = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
	LabelRequired = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED

	TypeMessage = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	Syntax      = "proto3"
)

type Dumper struct {
	fileDesc *descriptorpb.FileDescriptorProto
	ast      *ast.Container
}

func NewDumper(protoPackageName string) *Dumper {
	return &Dumper{
		fileDesc: &descriptorpb.FileDescriptorProto{Package: &protoPackageName, Syntax: &Syntax},
	}
}

func (d *Dumper) Dump(c *ast.Container) (*descriptorpb.FileDescriptorProto, error) {
	d.ast = c // 让 ast.Container 在方法之间共享
	messageDesc := &descriptorpb.DescriptorProto{
		Name: &c.Name,
	}

	d.fileDesc.MessageType = append(d.fileDesc.MessageType, messageDesc)

	for i, field := range c.Fields {
		fieldDesc, err := d.DumpField(field)
		if err != nil {
			return d.fileDesc, err
		}
		fieldNum := int32(i) + 1
		fieldDesc.Number = &fieldNum
		messageDesc.Field = append(messageDesc.Field, fieldDesc)

	}
	return d.fileDesc, nil
}

func (d *Dumper) DumpField(field *ast.Field) (*descriptorpb.FieldDescriptorProto, error) {
	if field.Basic != nil {
		return d.DumpBasicField(field.Basic)
	}
	if field.StructVector != nil {
		return d.DumpStructVectorField(field.StructVector)
	}
	if field.BasicVector != nil {
		return d.DumpBasicVectorField(field.BasicVector)
	}
	panic("invalid field")
}

func toType(b *ast.Basic) descriptorpb.FieldDescriptorProto_Type {
	if b.IsUINT32 {
		return descriptorpb.FieldDescriptorProto_TYPE_UINT32
	}
	if b.IsINT32 {
		return descriptorpb.FieldDescriptorProto_TYPE_INT32
	}
	if b.IsUINT64 {
		return descriptorpb.FieldDescriptorProto_TYPE_UINT64
	}
	if b.IsINT64 {
		return descriptorpb.FieldDescriptorProto_TYPE_INT64
	}
	if b.IsSTRING {
		return descriptorpb.FieldDescriptorProto_TYPE_STRING
	}
	if b.IsBOOL {
		return descriptorpb.FieldDescriptorProto_TYPE_BOOL
	}
	if b.IsDOUBLE {
		return descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	}
	if b.IsFLOAT {
		return descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	}
	if b.IsSTRING {
		return descriptorpb.FieldDescriptorProto_TYPE_STRING
	}
	if b.IsBOOL {
		return descriptorpb.FieldDescriptorProto_TYPE_BOOL
	}
	panic("invalid type")
}

func (d *Dumper) DumpBasicField(field *ast.BasicField) (*descriptorpb.FieldDescriptorProto, error) {
	typ := toType(&field.Type)
	return &descriptorpb.FieldDescriptorProto{Type: &typ, Name: &field.Name, Label: &LabelOptional}, nil
}

func (d *Dumper) DumpBasicVectorField(field *ast.BasicVectorField) (*descriptorpb.FieldDescriptorProto, error) {
	typ := toType(&field.Type)
	return &descriptorpb.FieldDescriptorProto{Type: &typ, Name: &field.Name, Label: &LabelRepeated}, nil
}

func (d *Dumper) DumpStructVectorField(field *ast.StructVectorField) (*descriptorpb.FieldDescriptorProto, error) {
	if len(field.StructList) == 0 {
		return nil, fmt.Errorf("field %s list<%s> has no element", field.Name, field.StructName)
	}
	messageDesc := &descriptorpb.DescriptorProto{
		Name: &field.Name,
	}

	d.fileDesc.MessageType = append(d.fileDesc.MessageType, messageDesc)

	for _, structInVector := range field.StructList {
		for i, subField := range structInVector.Fields {
			fieldDesc, err := d.DumpBasicField(subField)
			if err != nil {
				return nil, err
			}
			fieldNum := int32(i) + 1
			fieldDesc.Number = &fieldNum
			messageDesc.Field = append(messageDesc.Field, fieldDesc)
		}
	}
	return &descriptorpb.FieldDescriptorProto{Type: &TypeMessage, Name: &field.Name, Label: &LabelOptional, TypeName: &field.StructName}, nil
}
