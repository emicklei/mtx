package csv

import (
	std "encoding/csv"
	"errors"
	"io"
	"os"
)

func ScanSheet(filename string) (*Sheet, error) {
	s := NewSheet(filename)
	tab := s.Tab("main")

	f, _ := os.Open(filename)
	defer f.Close()
	r := std.NewReader(f)
	gotNames := false
	for {
		if gotNames {
			done := true
			for _, each := range tab.Columns {
				if each.ColumnType == UNKNOWN {
					done = false
				}
			}
			if done {
				// all types are known,no need to read more rows
				break
			}
		}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("unable to read CSV row")
		}
		if !gotNames {
			for _, each := range record {
				tab.Column(each).Type(UNKNOWN)
			}
			gotNames = true
		} else {
			// resolve unknown type
			for i, each := range record {
				typ := tab.Columns[i]
				if typ.GetDatatype() == UNKNOWN {
					// TODO can we detect nullable?
					typ.Type(DetectType(each))
				}
			}
		}
		//t.Log(record)
	}
	return s, nil
}
