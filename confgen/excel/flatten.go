package excel

import (
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
)

// 将 Excel 表拍平！
type Flattener struct {
	xlsxFile *xlsx.File
}

func NewFlattener(file *xlsx.File) *Flattener {
	return &Flattener{xlsxFile: file}
}

func (c *Flattener) Flatten(writer io.Writer, sheetName string) error {
	csvWriter := csv.NewWriter(writer)
	csvWriter.Comma = ' '
	defer csvWriter.Flush()
	sheet, ok := c.xlsxFile.Sheet[sheetName]
	if !ok {
		return fmt.Errorf("no sheet %s", sheetName)
	}
	for i := 0; i < sheet.MaxCol; i++ {
		if sheet.Cell(0, i).Value == "" {
			log.Printf("col %d empty, skip it", i)
			continue
		}
		if err := csvWriter.Write([]string{
			sheet.Cell(0, i).Value,
			sheet.Cell(1, i).Value,
			sheet.Cell(2, i).Value,
			fmt.Sprintf(`'%s'`, sheet.Cell(3, i).Value)},
		); err != nil {
			return err
		}
	}
	return nil
}
