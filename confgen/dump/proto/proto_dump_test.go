package proto

import (
	"github.com/stretchr/testify/assert"
	"github.com/xinnjie/confgen_by_parser/confgen/ast"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

func TestDumper_Dump(t *testing.T) {
	var (
		packageName           = "example.foo"
		syntax                = "proto3"
		containerName         = "Foo"
		fieldOneDesc          = "desc 1"
		fieldOneName          = "bar1"
		fieldOneNumber  int32 = 1
		fieldUint32Type       = descriptorpb.FieldDescriptorProto_TYPE_UINT32
		labelOptional         = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	)
	type args struct {
		c *ast.Container
	}
	tests := []struct {
		name    string
		args    args
		want    *descriptorpb.FileDescriptorProto
		wantErr bool
	}{
		{
			name: "basic field test",
			args: args{
				c: &ast.Container{
					Fields: []*ast.Field{
						{
							Basic: &ast.BasicField{
								Name: fieldOneName,
								Type: ast.Basic{
									IsUINT32: true,
								},
								Desc: fieldOneDesc,
							},
						},
					},
					Name: containerName,
				},
			},
			want: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: &containerName,
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   &fieldOneName,
								Number: &fieldOneNumber,
								Type:   &fieldUint32Type,
								Label:  &labelOptional,
							},
						},
					},
				},
				Package: &packageName,
				Syntax:  &syntax,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDumper(packageName)
			got, err := d.Dump(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
