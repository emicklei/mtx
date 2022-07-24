package csv

import (
	std "encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/emicklei/mtx"
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

type Option interface{}

func (s *Sheet) ToEntity(options ...Option) *mtx.Entity {
	tab := s.Tabs[0]
	pkg := mtx.NewPackage("sheet")
	ent := pkg.Entity(tab.Name)
	for _, each := range tab.Columns {
		mt := *each.ColumnType.AttributeDatatype
		a := ent.A(each.Name, mt, each.Documentation)
		if mt == mtx.UNKNOWN {
			// TODO helper func?
			a.Set("maperror", fmt.Sprintf("%s:%s", each.Name, each.ColumnType.Name))
		}
		a.IsNullable = each.IsNullable
	}
	return ent
}
