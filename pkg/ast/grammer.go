package ast

import (
	"github.com/alecthomas/participle/v2"
)

type Container struct {
	ID     string   `@Ident`
	Fields []*Field `@@+`
}

type Field struct {
	Scalar       *ScalarField       `(@@`
	ScalarVector *ScalarVectorField `|@@`
	StructVector *StructVectorField `|@@)`
}

type ScalarField struct {
	Name   string `@Ident`
	Scalar Scalar `@@`
	Desc   string `@String`
}

type StructVectorField struct {
	Name       string    `@Ident`
	StructName string    `Vector @Ident`
	Desc       string    `@String`
	StructList []*Struct `@@*`
}

type ScalarVectorField struct {
	Name       string    `@Ident`
	Scalar     Scalar    `Vector @@`
	Desc       string    `@String`
	StructList []*Struct `@@*`
}

type Scalar struct {
	IsEnum     bool `(@"E"`
	IsDateTime bool `| @"D")?` // 使用 string 存储
	IsUINT32   bool `( @Uint32`
	IsINT32    bool `| @Int32`
	IsUINT64   bool `| @Uint64`
	IsINT64    bool `| @Int64`
	IsSTRING   bool `| @StringT`
	IsBOOL     bool `| @Bool`
	IsDOUBLE   bool `| @Double`
	IsFLOAT    bool `| @Float)`
}

func (s *Scalar) Valid() bool {
	if s.IsEnum {
		if !s.IsInteger() {
			return false
		}
	}
	return true
}

func (s *Scalar) IsInteger() bool {
	if s.IsINT32 || s.IsUINT32 || s.IsINT64 || s.IsUINT64 {
		return true
	}
	return false
}

type Struct struct {
	Fields []*StructElement `LeftBracket @@* RightBracket`
}

type StructElement struct {
	Id   string `@Ident?`
	Type Scalar `@@`
	Desc string `@String`
}

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
