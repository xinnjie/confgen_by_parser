package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// for grammar test
func parseStructList(input string) (*StructInVector, error) {
	parser = participle.MustBuild(&StructInVector{}, parserOption...)
	symbol := &StructInVector{}
	err := parser.ParseString("struct parse test", input, symbol)
	return symbol, err
}

// for grammar test
func parseScalar(input string) (*Basic, error) {
	parser = participle.MustBuild(&Basic{}, parserOption...)
	symbol := &Basic{}
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

// for grammar test
func parseBasicInVector(input string) (*BasicInVector, error) {
	parser = participle.MustBuild(&BasicInVector{}, parserOption...)
	symbol := &BasicInVector{}
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
		want    *Basic
		wantErr bool
	}{
		{
			name: "uint32",
			args: args{
				input: "uint32",
			},
			want:    &Basic{IsUINT32: true},
			wantErr: false,
		},
		{
			name: "enum uint32",
			args: args{
				input: "enum uint32",
			},
			want: &Basic{
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
			want:    &Basic{IsINT32: true},
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
		want    *StructInVector
		wantErr bool
	}{
		{
			name: "struct",
			args: args{
				input: "[bar_1  uint32 '' \n" +
					"bar_2  uint32 '' ]",
			},
			want: &StructInVector{Fields: []*BasicField{
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
					Name: "foo_key",
					Type: Basic{IsUINT32: true},
					Desc: `'这是scalar字段'`,
				},
			},
			wantErr: false,
		},
		{
			name: "vector<int64> field",
			args: args{"foo vector int64 '这是vector<int64>字段' \n" +
				"[] int64 ''\n" +
				"[] int64 ''\n"},
			want: &Field{
				BasicVector: &BasicVectorField{
					Name: "foo",
					Type: Basic{IsINT64: true},
					Desc: `'这是vector<int64>字段'`,
					BasicList: &BasicInVector{
						Fields: []*BasicElement{
							{
								Type: Basic{IsINT64: true},
								Desc: `''`,
							},
							{
								Type: Basic{IsINT64: true},
								Desc: `''`,
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "vector<string> field",
			args: args{"foo vector string '这是vector<string>字段' \n" +
				"[] string ''\n" +
				"[] string ''\n"},
			want: &Field{
				BasicVector: &BasicVectorField{
					Name: "foo",
					Type: Basic{IsSTRING: true},
					Desc: `'这是vector<string>字段'`,
					BasicList: &BasicInVector{
						Fields: []*BasicElement{
							{
								Type: Basic{IsSTRING: true},
								Desc: `''`,
							},
							{
								Type: Basic{IsSTRING: true},
								Desc: `''`,
							},
						},
					},
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
							}},
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

func Test_parseBasicInVector(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *BasicInVector
		wantErr bool
	}{
		{
			name: "list of int64",
			args: args{
				input: "[] string ''\n" +
					"[] string ''\n",
			},
			want: &BasicInVector{
				Fields: []*BasicElement{
					{
						Type: Basic{IsSTRING: true},
						Desc: "''",
					},
					{
						Type: Basic{IsSTRING: true},
						Desc: "''",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBasicInVector(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBasicInVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
