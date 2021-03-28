package ast

type Container struct {
	Fields []*Field `parser:"@@+"`

	Name string
}

type Field struct {
	Basic        *BasicField        `parser:"(@@"`
	ScalarVector *ScalarVectorField `parser:"|@@"`
	StructVector *StructVectorField `parser:"|@@)"`
}

type BasicField struct {
	Scalar *ScalarField `parser:"(@@"`
	Bool   *BoolField   `parser:"|@@"`
	String *StringFiled `parser:"|@@)"`
}

type ScalarField struct {
	Name   string `parser:"@Ident"`
	Scalar Scalar `parser:"@@"`
	Desc   string `parser:"@String"`
}

type BoolField struct {
	Name string `parser:"@Ident Bool"`
	Desc string `parser:"@String"`
}

type StringFiled struct {
	Name string `parser:"@Ident StringT"`
	Desc string `parser:"@String"`
}

type StructVectorField struct {
	Name       string    `parser:"@Ident"`
	StructName string    `parser:"Vector @Ident"`
	Desc       string    `parser:"@String"`
	StructList []*Struct `parser:"@@*"`
}

type ScalarVectorField struct {
	Name       string    `parser:"@Ident"`
	Scalar     Scalar    `parser:"Vector @@"`
	Desc       string    `parser:"@String"`
	StructList []*Struct `parser:"@@*"`
}

type Scalar struct {
	IsEnum   bool `parser:"( @Enum"`
	IsTime   bool `parser:"| @Time)?"`
	IsUINT32 bool `parser:"( @Uint32"`
	IsINT32  bool `parser:"| @Int32"`
	IsUINT64 bool `parser:"| @Uint64"`
	IsINT64  bool `parser:"| @Int64"`
	IsSTRING bool `parser:"| @StringT"`
	IsBOOL   bool `parser:"| @Bool"`
	IsDOUBLE bool `parser:"| @Double"`
	IsFLOAT  bool `parser:"| @Float)"`
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
	Fields []*StructElement `parser:"LeftBracket @@* RightBracket"`
}

type StructElement struct {
	Id   string `parser:"@Ident?"`
	Type Scalar `parser:"@@"`
	Desc string `parser:"@String"`
}
