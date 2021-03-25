package excel

import (
	"bytes"
	"github.com/tealeg/xlsx"
	"github.com/xinnjie/confgen_by_parser/pkg/ast"
	"log"
)

type Loader struct {
	xlsxName string
	exporter *Exporter
}

func NewLoader(xlsxName string) (*Loader, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxName)
	if err != nil {
		log.Print("open xlsx filed", err)
		return nil, err
	}
	exporter := NewExporter(xlsxFile)
	return &Loader{xlsxName: xlsxName, exporter: exporter}, nil
}

func (l *Loader) Load(sheetName string) error {
	buf := &bytes.Buffer{}
	if err := l.exporter.Export(buf, sheetName); err != nil {
		return err
	}
	c, err := ast.GenAst(buf, l.xlsxName)
	if err != nil {
		return err
	}
	c.Name = sheetName
	return nil
}
