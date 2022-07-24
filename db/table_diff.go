package db

import "github.com/emicklei/mtx"

type TableDiff struct {
	ColumnAdditions []*Column
	ColumnChanges   []*Column
	// same size as ColumnChanges, each entry is a list of change aspects found
	// such as name, description, type etc.
	ChangeAspects  [][]string
	ColumnRemovals []*Column
}

func (t *Table) Diff(other *Table) TableDiff {
	diff := TableDiff{}
	// columns not in other or different
	for _, left := range t.Columns {
		right, ok := mtx.FindByName(other.Columns, left.Name)
		if !ok { // not in other
			diff.ColumnRemovals = append(diff.ColumnRemovals, left)
		} else {
			aspects := []string{}
			// also in other, may have changes
			if left.ColumnType.Name != right.ColumnType.Name {
				aspects = append(aspects, "name")
				diff.ColumnChanges = append(diff.ColumnChanges, right)
			}
			if left.IsPrimary != right.IsPrimary {
				aspects = append(aspects, "isprimary")
				diff.ColumnChanges = append(diff.ColumnChanges, right)
			}
			if left.IsNullable != right.IsNullable {
				aspects = append(aspects, "nullable")
				diff.ColumnChanges = append(diff.ColumnChanges, right)
			}
			diff.ChangeAspects = append(diff.ChangeAspects, aspects)
		}
	}
	// columns not in t
	for _, right := range other.Columns {
		_, ok := mtx.FindByName(t.Columns, right.Name)
		if !ok { // not in t
			diff.ColumnAdditions = append(diff.ColumnAdditions, right)
		}
	}
	return diff
}
