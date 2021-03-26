package excel

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestCsvExport_ExportCsv(t *testing.T) {
	xlsxFile, err := xlsx.OpenFile("../../testdata/basic-field.xlsx")
	assert.NoError(t, err)
	exporter := NewFlattener(xlsxFile)
	buf := &bytes.Buffer{}
	assert.NoError(t, exporter.Flatten(buf, "Basic"))
	fmt.Println(buf.String())
	assert.EqualValues(t, "title  string '标题'\nseason  uint32 '赛季编号'\n", buf.String())
}
