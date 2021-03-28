package load

import (
	"bytes"
	"github.com/xinnjie/confgen_by_parser/confgen/dump/proto"
	"io"
	"log"

	"github.com/tealeg/xlsx/v3"
	"github.com/xinnjie/confgen_by_parser/confgen/ast"
)

type Loader struct {
	xlsxName  string
	flattener *Flattener
	ast       *ast.Container
}

func NewLoader(xlsxName string) (*Loader, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxName)
	if err != nil {
		log.Print("open xlsx filed", err)
		return nil, err
	}
	exporter := NewFlattener(xlsxFile)
	return &Loader{xlsxName: xlsxName, flattener: exporter}, nil
}

func (l *Loader) Load(sheetName string) error {
	buf := &bytes.Buffer{}
	if err := l.flattener.Flatten(buf, sheetName); err != nil {
		return err
	}
	c, err := ast.GenAst(buf, l.xlsxName)
	if err != nil {
		return err
	}
	c.Name = sheetName
	l.ast = c
	return nil
}

func (l *Loader) OutputProto(w io.Writer, protoPackageName string) error {
	dumper := proto.NewDumper(protoPackageName)
	fileProtoDesc, err := dumper.Dump(l.ast)
	if err != nil {
		return err
	}
	proto.Restore(w, fileProtoDesc)
	return nil
}
