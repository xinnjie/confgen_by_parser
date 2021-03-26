package ast

type Container struct {
	Fields []*Field `@@+`

	Name string
}

type Field struct {
	Basic        *BasicField        `(@@`
	ScalarVector *ScalarVectorField `|@@`
	StructVector *StructVectorField `|@@)`
}

type BasicField struct {
	Scalar *ScalarField `(@@`
	Bool   *BoolField   `|@@`
	String *StringFiled `|@@)`
}

type ScalarField struct {
	Name   string `@Ident`
	Scalar Scalar `@@`
	Desc   string `@String`

	Value string
}

type BoolField struct {
	Name string `@Ident Bool`
	Desc string `@String`

	Value string
}

type StringFiled struct {
	Name string `@Ident StringT`
	Desc string `@String`

	Value string
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
