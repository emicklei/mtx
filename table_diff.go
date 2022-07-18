package mtx

type TableDiff[C ExtendsColumn, D ExtendsDatatype] struct {
	ColumnAdditions []*Column[C, D]
	ColumnChanges   []*Column[C, D]
	ColumnRemovals  []*Column[C, D]
}

func (t *Table[T, C, D]) Diff(other *Table[T, C, D]) TableDiff[C, D] {
	diff := TableDiff[C, D]{}
	// columns not in other or different
	for _, left := range t.Columns {
		right, ok := FindByName(other.Columns, left.Name)
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
					if left.IsNotNull != left.IsNotNull {
						diff.ColumnChanges = append(diff.ColumnChanges, right)
					}
				}
			}
		}
	}
	// columns not in t
	for _, right := range other.Columns {
		_, ok := FindByName(t.Columns, right.Name)
		if !ok { // not in t
			diff.ColumnAdditions = append(diff.ColumnAdditions, right)
		}
	}
	return diff
}
