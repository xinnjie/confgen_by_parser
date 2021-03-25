package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
)

// 摊平 Excel 表格
type Exporter struct {
	xlsxName string
	xlsxFile *xlsx.File
}

func NewExporter(xlsxName string) (*Exporter, error) {
	xlsxFile, err := xlsx.OpenFile(xlsxName)
	if err != nil {
		log.Print("open xlsx filed", err)
		return nil, err
	}
	return &Exporter{xlsxName: xlsxName, xlsxFile: xlsxFile}, nil
}

func (c *Exporter) Export(writer io.Writer, sheetName string) error {
	sheet, ok := c.xlsxFile.Sheet[sheetName]
	if !ok {
		return fmt.Errorf("no sheet %s in %s", sheetName, c.xlsxName)
	}
	for i := 0; i < sheet.MaxCol; i++ {
		if sheet.Cell(0, i).Value == "" {
			log.Printf("col %d empty skip", i)
			continue
		}
		_, err := fmt.Fprintf(writer, "%s %s %s \"%s\"\n", sheet.Cell(0, i).Value,
			sheet.Cell(1, i).Value,
			sheet.Cell(2, i).Value,
			sheet.Cell(3, i).Value)
		if err != nil {
			return err
		}
	}
	return nil
}
