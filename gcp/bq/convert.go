package bq

import (
	"encoding/json"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

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

func ToTable(ent *mtx.Entity) *db.Table {
	ds := NewDataset(ent.Name)
	tab := ds.Table(ent.Name)
	// TODO who is primary
	for _, each := range ent.Attributes {
		mt := MappedAttributeType(each.AttributeType)
		c := tab.C(each.Name, mt, each.Documentation)
		if mt == mtx.UNKNOWN {
			c.Set("maperror", each.AttributeType.Name)
		}
		c.IsNullable = each.IsNullable
	}
	return tab
}
