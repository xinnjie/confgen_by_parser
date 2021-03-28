package load

import (
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"io"
	"log"
	"strings"
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
		cellRow0, err := sheet.Cell(0, i)
		if err != nil {
			// not possible
			panic(err)
		}
		cellRow1, err := sheet.Cell(1, i)
		if err != nil {
			// not possible
			panic(err)
		}
		cellRow2, err := sheet.Cell(2, i)
		if err != nil {
			// not possible
			panic(err)
		}
		cellRow3, err := sheet.Cell(3, i)
		if err != nil {
			// not possible
			panic(err)
		}
		// 跳过空列
		if cellRow0.Value == "" {
			log.Printf("col %d empty, skip it", i)
			continue
		}

		cellRow3.Value = fmt.Sprintf(`'%s'`, cellRow3.Value)
		// 结构体变换结构
		structEnd := len(cellRow0.Value) > 2 && strings.HasSuffix(cellRow0.Value, "]")
		if structEnd {
			cellRow0.Value = cellRow0.Value[:len(cellRow0.Value)-1]
			cellRow3.Value = cellRow3.Value + "]"
		}

		if err := csvWriter.Write([]string{
			cellRow0.Value,
			cellRow1.Value,
			cellRow2.Value,
			cellRow3.Value},
		); err != nil {
			return err
		}
	}
	return nil
}
