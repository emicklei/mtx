package mtx

type DatabaseView struct {
	*Named
	ColumnExpressions []any
}
