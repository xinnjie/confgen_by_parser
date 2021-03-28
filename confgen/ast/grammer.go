package ast

type Container struct {
	Fields []*Field `parser:"@@+"`

	Name string
}

type Field struct {
	Basic        *BasicField        `parser:"(@@"`
	BasicVector  *BasicVectorField  `parser:"|@@"`
	StructVector *StructVectorField `parser:"|@@)"`
}

type BasicField struct {
	Name string `parser:"@Ident"`
	Type Basic  `parser:"@@"`
	Desc string `parser:"@String"`
}

type StructVectorField struct {
	Name       string            `parser:"@Ident"`
	StructName string            `parser:"Vector @Ident"`
	Desc       string            `parser:"@String"`
	StructList []*StructInVector `parser:"@@*"`
}

type BasicVectorField struct {
	Name      string         `parser:"@Ident"`
	Type      Basic          `parser:"Vector @@"`
	Desc      string         `parser:"@String"`
	BasicList *BasicInVector `parser:"@@*"`
}

type Basic struct {
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

func (s *Basic) Valid() bool {
	if s.IsEnum {
		if !s.IsInteger() {
			return false
		}
	}
	return true
}

func (s *Basic) IsInteger() bool {
	if s.IsINT32 || s.IsUINT32 || s.IsINT64 || s.IsUINT64 {
		return true
	}
	return false
}

type BasicInVector struct {
	Fields []*BasicElement `parser:"@@+"` // TODO @@* 会死循环, parser库可以改进
}

type BasicElement struct {
	Type Basic  `parser:"LeftBracket RightBracket @@"`
	Desc string `parser:"@String"`
}

type StructInVector struct {
	Fields []*BasicField `parser:"LeftBracket @@* RightBracket"`
}
