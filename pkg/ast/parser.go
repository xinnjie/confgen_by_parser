package ast

import "github.com/alecthomas/participle/v2"

var (
	parserOption = []participle.Option{
		participle.UseLookahead(2),
		participle.Lexer(tokens),
		participle.Elide("Whitespace"),
	}

	parser = participle.MustBuild(&Map{}, parserOption...)
)
