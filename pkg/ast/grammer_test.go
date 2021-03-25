package ast

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

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
				input: "E uint32",
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

func Test_parseEnumField(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *EnumElement
		wantErr bool
	}{
		{
			name: "valid enum",
			args: args{input: "[实例类型]实例类型1 1 E_BAR_1 \n"},
			want: &EnumElement{
				EnumLiteral: "[实例类型]实例类型1",
				EnumValue:   1,
				ID:          "E_BAR_1",
			},
			wantErr: false,
		},
		{
			name:    "invalid enum",
			args:    args{input: "实例类型实例类型1 1 E_BAR_1 \n"},
			want:    &EnumElement{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEnumField(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEnumField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnumField() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseEnum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *Enum
		wantErr bool
	}{
		{
			name: "enum without field",
			args: args{
				input: "{BarEnum} \n",
			},
			want: &Enum{
				EnumType: "BarEnum",
			},
			wantErr: false,
		},
		{
			name: "enum with field",
			args: args{
				input: "{BarEnum} \n" +
					"[实例类型]实例类型1 1 E_BAR_1 \n" +
					"[实例类型]实例类型2 2 E_BAR_2 \n",
			},
			want: &Enum{
				EnumType: "BarEnum",
				EnumElms: []*EnumElement{
					{
						EnumLiteral: "[实例类型]实例类型1",
						EnumValue:   1,
						ID:          "E_BAR_1",
					},
					{
						EnumLiteral: "[实例类型]实例类型2",
						EnumValue:   2,
						ID:          "E_BAR_2",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEnum(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEnum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnum() got = %v, want %v", got, tt.want)
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
				input: "{bar_1  uint32 \"\" \n" +
					"bar_2  uint32 \"\" }",
			},
			want: &Struct{Fields: []*StructElement{
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
			args: args{"foo_key  uint32 \"这是scalar字段\""},
			want: &Field{
				Scalar: &ScalarField{
					Name:   "foo_key",
					Scalar: Scalar{IsUINT32: true},
					Desc:   `"这是scalar字段"`,
				},
			},
			wantErr: false,
		},
		{
			name: "vector<int64> field",
			args: args{"foo vector int64 \"这是vector<int64>字段\" \n" +
				"{int64 \"\"}\n" +
				"{int64 \"\"}\n"},
			want: &Field{
				ScalarVector: &ScalarVectorField{
					Name:   "foo",
					Scalar: Scalar{IsINT64: true},
					Desc:   `"这是vector<int64>字段"`,
					StructList: []*Struct{
						{
							Fields: []*StructElement{
								{
									Id:   "",
									Type: Scalar{IsINT64: true},
									Desc: `""`,
								},
							},
						},
						{
							Fields: []*StructElement{
								{
									Id:   "",
									Type: Scalar{IsINT64: true},
									Desc: `""`,
								},
							},
						}},
				},
			},
			wantErr: false,
		},
		{
			name: "vector<struct> field",
			args: args{"foo vector FooStruct \"这是vector<struct>字段\" \n" +
				"{bar_1  uint32 \"\" \n" +
				"bar_2  uint32 \"\" }\n" +
				"{bar_1  uint32 \"\" \n" +
				"bar_2  uint32 \"\" }\n"},
			want: &Field{
				StructVector: &StructVectorField{
					Name:       "foo",
					StructName: "FooStruct",
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
							}},
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
