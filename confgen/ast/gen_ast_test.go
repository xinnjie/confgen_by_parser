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
						Basic: &BasicField{
							Name: "foo_key",
							Type: Basic{IsUINT32: true},
							Desc: `'这是scalar字段'`,
						},
					},
					{
						Basic: &BasicField{
							Name: "foo_uint32",
							Type: Basic{IsUINT32: true},
							Desc: `''`,
						},
					},
					{
						StructVector: &StructVectorField{
							Name:       "bar",
							StructName: "BarStruct",
							Desc:       `'这是vector<struct>字段'`,
							StructList: []*StructInVector{
								{
									Fields: []*BasicField{
										{
											Name: "bar_1",
											Type: Basic{IsUINT32: true},
											Desc: `''`,
										},
										{
											Name: "bar_2",
											Type: Basic{IsUINT32: true},
											Desc: `''`,
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
