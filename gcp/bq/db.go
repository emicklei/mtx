package bq

import (
	"encoding/json"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

func NewDataset(name string) *db.Database {
	return &db.Database{
		Named:      mtx.N("spanner.Dataset", name),
		Extensions: new(DatabaseExtensions),
	}
}

func ToJSONSchema(tab *db.Table) string {
	cols := []JSONSchemaColumn{}
	for _, each := range tab.Columns {
		col := JSONSchemaColumn{
			Description: each.Documentation,
			Name:        each.Name,
			Type:        each.ColumnType.Name,
			Mode:        "NULLABLE",
		}
		if !each.IsNullable {
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
