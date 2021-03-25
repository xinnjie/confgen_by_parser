package ast

import (
	"io"
	"log"
)

func GenAst(defineReader io.Reader, optionalFileName string) (*Container, error) {
	ini := &Container{}
	err := parser.Parse(optionalFileName, defineReader, ini)
	if err != nil {
		log.Print("parse err", err)
		return nil, err
	}
	return ini, nil
}
