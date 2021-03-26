package ast

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenAst(t *testing.T) {
	type args struct {
		defineString string
	}
	tests := []struct {
		name string
		args args
		want *Container
	}{
		{
			"one Key map parser test ",
			args{
				defineString: "foo_key  uint32 '这是scalar字段' \n" +
					"foo_uint32  uint32 '' \n" +
					"bar vector BarStruct '这是vector<struct>字段' \n" +
					"[bar_1  uint32 '' \n" +
					"bar_2  uint32 '' ]\n",
			},
			&Container{
				Fields: []*Field{
					{
						Scalar: &ScalarField{
							Name:   "foo_key",
							Scalar: Scalar{IsUINT32: true},
							Desc:   `"这是scalar字段"`,
						},
					},
					{
						Scalar: &ScalarField{
							Name:   "foo_uint32",
							Scalar: Scalar{IsUINT32: true},
							Desc:   `""`,
						},
					},
					{
						StructVector: &StructVectorField{
							Name:       "bar",
							StructName: "BarStruct",
							Desc:       `"这是vector<struct>字段"`,
							StructList: []*Struct{
								{
									Fields: []*StructElement{
										{
											Id:   "bar_1",
											Type: Scalar{IsUINT32: true},
											Desc: `""`,
										},
										{
											Id:   "bar_2",
											Type: Scalar{IsUINT32: true},
											Desc: `""`,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defineString := bytes.NewBufferString(tt.args.defineString)

			got, err := GenAst(defineString, tt.name)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
