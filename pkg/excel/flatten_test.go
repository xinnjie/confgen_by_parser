package excel

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsvExport_ExportCsv(t *testing.T) {
	exporter, err := NewExporter("../../testdata/basic-field.xlsx")
	assert.NoError(t, err)
	buf := &bytes.Buffer{}
	assert.NoError(t, exporter.Export(buf, "Basic"))
	fmt.Println(buf.String())
	assert.EqualValues(t, "title  string \"标题\"\nseason  uint32 \"赛季编号\"\n", buf.String())
}
