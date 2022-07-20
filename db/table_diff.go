package db

import "github.com/emicklei/mtx"

type TableDiff struct {
	ColumnAdditions []*Column
	ColumnChanges   []*Column
	ColumnRemovals  []*Column
}

func (t *Table) Diff(other *Table) TableDiff {
	diff := TableDiff{}
	// columns not in other or different
	for _, left := range t.Columns {
		right, ok := mtx.FindByName(other.Columns, left.Name)
		if !ok { // not in other
			diff.ColumnRemovals = append(diff.ColumnRemovals, left)
		} else {
			// also in other, may have changes
			if left.ColumnType.Name != right.ColumnType.Name {
				diff.ColumnChanges = append(diff.ColumnChanges, right)
			} else {
				if left.IsPrimary != left.IsPrimary {
					diff.ColumnChanges = append(diff.ColumnChanges, right)
				} else {
					if left.IsNullable != left.IsNullable {
						diff.ColumnChanges = append(diff.ColumnChanges, right)
					}
				}
			}
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
