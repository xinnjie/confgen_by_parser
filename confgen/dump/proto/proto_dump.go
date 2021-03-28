package proto

import (
	"github.com/xinnjie/confgen_by_parser/confgen/ast"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	LabelOptional = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
)

type Dumper struct {
	fileDesc *descriptorpb.FileDescriptorProto
	ast      *ast.Container
}

func NewDumper() *Dumper {
	return &Dumper{
		fileDesc: &descriptorpb.FileDescriptorProto{},
	}
}

func (d *Dumper) Dump(c *ast.Container) (*descriptorpb.FileDescriptorProto, error) {
	d.ast = c // 让 ast.Container 在方法之间共享
	messageDesc := &descriptorpb.DescriptorProto{
		Name: &c.Name,
	}
	defer func() {
		d.fileDesc.MessageType = append(d.fileDesc.MessageType, messageDesc)
	}()
	for i, field := range c.Fields {
		fieldDesc, err := d.DumpField(field, messageDesc)
		if err != nil {
			return d.fileDesc, err
		}
		fieldNum := int32(i) + 1
		fieldDesc.Number = &fieldNum
		messageDesc.Field = append(messageDesc.Field, fieldDesc)

	}
	return d.fileDesc, nil
}

func (d *Dumper) DumpField(field *ast.Field, msgDesc *descriptorpb.DescriptorProto) (*descriptorpb.FieldDescriptorProto, error) {
	if field.Basic != nil {
		return d.DumpBasicField(field.Basic)
	}
	panic("invalid field")
}

func (d *Dumper) DumpBasicField(basic *ast.BasicField) (*descriptorpb.FieldDescriptorProto, error) {
	if basic.Scalar != nil {
		return d.DumpScalarField(basic.Scalar)
	}
	if basic.String != nil {
		return d.DumpStringField(basic.String)
	}
	return d.DumpBoolField(basic.Bool)
}

func (d *Dumper) DumpScalarField(scalar *ast.ScalarField) (*descriptorpb.FieldDescriptorProto, error) {
	typ := func() descriptorpb.FieldDescriptorProto_Type {
		if scalar.Scalar.IsUINT32 {
			return descriptorpb.FieldDescriptorProto_TYPE_UINT32
		}
		if scalar.Scalar.IsINT32 {
			return descriptorpb.FieldDescriptorProto_TYPE_INT32
		}
		if scalar.Scalar.IsUINT64 {
			return descriptorpb.FieldDescriptorProto_TYPE_UINT64
		}
		if scalar.Scalar.IsINT64 {
			return descriptorpb.FieldDescriptorProto_TYPE_INT64
		}
		if scalar.Scalar.IsSTRING {
			return descriptorpb.FieldDescriptorProto_TYPE_STRING
		}
		if scalar.Scalar.IsBOOL {
			return descriptorpb.FieldDescriptorProto_TYPE_BOOL
		}
		if scalar.Scalar.IsDOUBLE {
			return descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
		}
		if scalar.Scalar.IsFLOAT {
			return descriptorpb.FieldDescriptorProto_TYPE_FLOAT
		}
		panic("invalid type")
	}()
	return &descriptorpb.FieldDescriptorProto{Type: &typ, Name: &scalar.Name, Label: &LabelOptional}, nil
}

func (d *Dumper) DumpStringField(field *ast.StringFiled) (*descriptorpb.FieldDescriptorProto, error) {
	typ := descriptorpb.FieldDescriptorProto_TYPE_STRING
	return &descriptorpb.FieldDescriptorProto{Type: &typ, Name: &field.Name, Label: &LabelOptional}, nil
}

func (d *Dumper) DumpBoolField(field *ast.BoolField) (*descriptorpb.FieldDescriptorProto, error) {
	typ := descriptorpb.FieldDescriptorProto_TYPE_BOOL
	return &descriptorpb.FieldDescriptorProto{Type: &typ, Name: &field.Name, Label: &LabelOptional}, nil
}
