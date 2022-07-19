package bq

import (
	"encoding/json"

	"github.com/emicklei/mtx"
)

func NewDataset(name string) *mtx.Database {
	return &mtx.Database{
		Named:      mtx.N("spanner.Dataset", name),
		Extensions: new(DatabaseExtensions),
	}
}

func ToJSONSchema(tab *mtx.Table) string {
	cols := []JSONSchemaColumn{}
	for _, each := range tab.Columns {
		col := JSONSchemaColumn{
			Description: each.Documentation,
			Name:        each.Name,
			Type:        each.ColumnType.Name,
			Mode:        "NULLABLE",
		}
		if each.IsNotNull {
			col.Mode = "REQUIRED"
		}
		cols = append(cols, col)
	}
	data, _ := json.MarshalIndent(cols, "", "  ")
	return string(data)
}

type JSONSchemaColumn struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Mode        string `json:"mode"`
}
