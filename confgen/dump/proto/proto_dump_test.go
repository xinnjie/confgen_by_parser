package proto

import (
	"github.com/stretchr/testify/assert"
	"github.com/xinnjie/confgen_by_parser/confgen/ast"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

func TestDumper_Dump(t *testing.T) {
	var (
		ContainerName         = "Foo"
		FieldOneDesc          = "desc 1"
		FieldOneName          = "bar1"
		FieldOneNumber  int32 = 1
		FieldUint32Type       = descriptorpb.FieldDescriptorProto_TYPE_UINT32
		LabelOptional         = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
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
								Scalar: &ast.ScalarField{
									Name: FieldOneName,
									Scalar: ast.Scalar{
										IsUINT32: true,
									},
									Desc: FieldOneDesc,
								},
							},
						},
					},
					Name: ContainerName,
				},
			},
			want: &descriptorpb.FileDescriptorProto{
				MessageType: []*descriptorpb.DescriptorProto{
					{
						Name: &ContainerName,
						Field: []*descriptorpb.FieldDescriptorProto{
							{
								Name:   &FieldOneName,
								Number: &FieldOneNumber,
								Type:   &FieldUint32Type,
								Label:  &LabelOptional,
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDumper()
			got, err := d.Dump(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
