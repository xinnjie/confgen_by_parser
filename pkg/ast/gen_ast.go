package ast

import (
	"io"
)

func GenAst(defineReader io.Reader, optionalFileName string) Map {
	ini := &Map{}
	err := parser.Parse(optionalFileName, defineReader, ini)
	if err != nil {
		panic(err)
	}
	return *ini
}
