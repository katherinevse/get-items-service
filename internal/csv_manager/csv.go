package csv_manager

import (
	"encoding/csv"
	"io"
	"os"
)

type CsvManager struct {
}

func (m *CsvManager) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (m *CsvManager) NewReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}
