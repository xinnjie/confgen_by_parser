package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
)

// 摊平 Excel 表格
type Exporter struct {
	xlsxFile *xlsx.File
}

func NewExporter(file *xlsx.File) *Exporter {
	return &Exporter{xlsxFile: file}
}

func (c *Exporter) Export(writer io.Writer, sheetName string) error {
	sheet, ok := c.xlsxFile.Sheet[sheetName]
	if !ok {
		return fmt.Errorf("no sheet %s", sheetName)
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
