package proto

import (
	"fmt"
	"google.golang.org/protobuf/types/descriptorpb"
	"io"
)

func Restore(w io.Writer, proto *descriptorpb.FileDescriptorProto) {
	fmt.Fprintf(w, "syntax = \"%s\";\n\n"+
		"package %s;\n\n", proto.GetSyntax(), proto.GetPackage())

	for i, message := range proto.MessageType {
		RestoreFromMessage(w, message)
		if i != len(proto.MessageType)-1 {
			fmt.Fprintf(w, "\n")
		}
	}
}

func RestoreFromMessage(w io.Writer, msg *descriptorpb.DescriptorProto) {
	fmt.Fprintf(w, "message %s {\n", *msg.Name)
	for _, field := range msg.Field {
		RestoreFromField(w, field)
	}
	fmt.Fprintf(w, "}\n")
}

func RestoreFromField(w io.Writer, field *descriptorpb.FieldDescriptorProto) {
	fmt.Fprintf(w, "  %s%s %s = %d;\n", labelString(field.GetLabel()), typeString(field), *field.Name, *field.Number)
}

func typeString(field *descriptorpb.FieldDescriptorProto) string {
	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return "double"
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return "float"
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		return "fixed64"
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		return "fixed32"
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return "bytes"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		return "sfixed32"
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		return "sfixed64"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		return "sint32"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return "sint64"
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		fallthrough
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		return field.GetTypeName()
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		panic("not implemented")
	}
	panic("not implemented")
}

func labelString(label descriptorpb.FieldDescriptorProto_Label) string {
	switch label {
	case descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL:
		// TODO 按照 syntax 选择是否打印 optional
		return ""
	case descriptorpb.FieldDescriptorProto_LABEL_REQUIRED:
		return "required "
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		return "repeated "
	}
	panic("invalid label")
}
