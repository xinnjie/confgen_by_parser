package ast

import (
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

var (
	tokens = stateful.MustSimple([]stateful.Rule{
		// 关键字
		{"LeftBracket", `\[`, nil},
		{"RightBracket", `\]`, nil},
		{"Vector", `vector`, nil},
		{"DOLLAR_SPLIT", `\$`, nil},
		{"Uint32", "uint32", nil},
		{"Int32", "int32", nil},
		{"Uint64", "uint64", nil},
		{"Int64", "int64", nil},
		{"Bool", "bool", nil},
		{"Double", "double", nil},
		{"Float", "float", nil},
		{"StringT", "string", nil},

		{"String", `'[^']*'`, nil},
		{"Number", `[-+]?(\d*\.)?\d+`, nil},
		{"Ident", `[a-zA-Z_]\w*`, nil},
		{"Whitespace", `[ \t\n\r]+`, nil},
	})
)
