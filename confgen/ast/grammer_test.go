package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// for grammar test
func parseStructList(input string) (*Struct, error) {
	parser = participle.MustBuild(&Struct{}, parserOption...)
	symbol := &Struct{}
	err := parser.ParseString("struct parse test", input, symbol)
	return symbol, err
}

// for grammar test
func parseScalar(input string) (*Scalar, error) {
	parser = participle.MustBuild(&Scalar{}, parserOption...)
	symbol := &Scalar{}
	err := parser.ParseString("scalar parse test", input, symbol)
	return symbol, err
}

// for grammar test
func parseField(input string) (*Field, error) {
	parser = participle.MustBuild(&Field{}, parserOption...)
	symbol := &Field{}
	err := parser.ParseString("filed parse test", input, symbol)
	return symbol, err
}

func Test_parseScalar(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *Scalar
		wantErr bool
	}{
		{
			name: "uint32",
			args: args{
				input: "uint32",
			},
			want:    &Scalar{IsUINT32: true},
			wantErr: false,
		},
		{
			name: "enum uint32",
			args: args{
				input: "enum uint32",
			},
			want: &Scalar{
				IsEnum:   true,
				IsUINT32: true,
			},
			wantErr: false,
		},
		{
			name: "int32",
			args: args{
				input: "int32",
			},
			want:    &Scalar{IsINT32: true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseScalar(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseScalar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseScalar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseStructList(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *Struct
		wantErr bool
	}{
		{
			name: "struct",
			args: args{
				input: "[bar_1  uint32 '' \n" +
					"bar_2  uint32 '' ]",
			},
			want: &Struct{Fields: []*StructElement{
				{
					Id:   "bar_1",
					Type: Scalar{IsUINT32: true},
					Desc: `''`,
				},
				{
					Id:   "bar_2",
					Type: Scalar{IsUINT32: true},
					Desc: `''`,
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStructList(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStructList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func Test_parseField(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *Field
		wantErr bool
	}{
		{
			name: "scalar field",
			args: args{"foo_key  uint32 '这是scalar字段'"},
			want: &Field{
				Basic: &BasicField{
					Scalar: &ScalarField{
						Name:   "foo_key",
						Scalar: Scalar{IsUINT32: true},
						Desc:   `'这是scalar字段'`,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "vector<int64> field",
			args: args{"foo vector int64 '这是vector<int64>字段' \n" +
				"[int64 '']\n" +
				"[int64 '']\n"},
			want: &Field{
				ScalarVector: &ScalarVectorField{
					Name:   "foo",
					Scalar: Scalar{IsINT64: true},
					Desc:   `'这是vector<int64>字段'`,
					StructList: []*Struct{
						{
							Fields: []*StructElement{
								{
									Id:   "",
									Type: Scalar{IsINT64: true},
									Desc: `''`,
								},
							},
						},
						{
							Fields: []*StructElement{
								{
									Id:   "",
									Type: Scalar{IsINT64: true},
									Desc: `''`,
								},
							},
						}},
				},
			},
			wantErr: false,
		},
		{
			name: "vector<struct> field",
			args: args{"foo vector FooStruct '这是vector<struct>字段' \n" +
				"[bar_1  uint32 '' \n" +
				"bar_2  uint32 '' ]\n" +
				"[bar_1  uint32 '' \n" +
				"bar_2  uint32 '' ]\n"},
			want: &Field{
				StructVector: &StructVectorField{
					Name:       "foo",
					StructName: "FooStruct",
					Desc:       `'这是vector<struct>字段'`,
					StructList: []*Struct{
						{
							Fields: []*StructElement{
								{
									Id:   "bar_1",
									Type: Scalar{IsUINT32: true},
									Desc: `''`,
								},
								{
									Id:   "bar_2",
									Type: Scalar{IsUINT32: true},
									Desc: `''`,
								},
							}},
						{
							Fields: []*StructElement{
								{
									Id:   "bar_1",
									Type: Scalar{IsUINT32: true},
									Desc: `''`,
								},
								{
									Id:   "bar_2",
									Type: Scalar{IsUINT32: true},
									Desc: `''`,
								},
							}},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseField(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
