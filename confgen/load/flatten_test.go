package load

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx/v3"
	"testing"
)

func TestCsvExport_ExportCsv(t *testing.T) {
	xlsxFile, err := xlsx.OpenFile("../../testdata/type-field.xlsx")
	assert.NoError(t, err)
	exporter := NewFlattener(xlsxFile)
	buf := &bytes.Buffer{}
	assert.NoError(t, exporter.Flatten(buf, "DifferentTypes"))
	fmt.Println(buf.String())
	assert.EqualValues(t, `foo_string  string '注释行：这是一串字符'
foo_uint32  uint32 '这是32位无符号整型'
foo_int32  int32 '这是32位有符号整型'
foo_uint64  uint64 '64位无符号整型'
foo_int64  int64 '64位有符号整型'
foo_list_string vector string '字符串数组'
[]  string '字符串数组元素'
[]  string '字符串数组元素'
foo_list_int64 vector int64 'int数组'
[]  int64 'int数组元素'
[]  int64 'int数组元素'
foo_list_struct vector BarStruct '结构体数组'
[id  int64 '结构体数组元素'
cnt  int64 '结构体数组元素']
[id  int64 '结构体数组元素'
cnt  int64 '结构体数组元素']
`, buf.String())
}
