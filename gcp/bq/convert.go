package bq

import (
	"encoding/json"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/emicklei/mtx/db"
)

// ToJSONSchema returns a BigQuery Schema in JSON string.
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
		if each.ColumnType.Name == "RECORD" && len(Extensions(each).NestedColumns) > 0 {
			for _, other := range Extensions(each).NestedColumns {
				// TODO recurs
				nestedcol := JSONSchemaColumn{
					Description: other.Documentation,
					Name:        other.Name,
					Type:        other.ColumnType.Name,
					Mode:        "NULLABLE",
				}
				if !other.IsNullable {
					nestedcol.Mode = "REQUIRED"
				}
				col.Fields = append(col.Fields, nestedcol)
			}
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
	// for record
	Fields []JSONSchemaColumn `json:"fields,omitempty"`
	Mode   string             `json:"mode"`
}

func ToTable(ent *basic.Entity) *db.Table {
	ds := NewDataset(ent.Name)
	tab := ds.Table(ent.Name)
	// TODO who is primary
	for _, each := range ent.Attributes {
		mt := MappedAttributeType(each.AttributeType)
		c := tab.C(each.Name, mt, each.Documentation)
		if mt.Equal(mtx.Unknown) {
			c.Set("maperror", each.AttributeType.Name)
		}
	}
	return tab
}

// ToBasicType returns a mapped basic Datatype
func ToBasicType(dt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(dt, registry.Class())

	bt := *dt.BasicDatatype
	bt.CopyPropertiesFrom(dt.Named)
	if dt.IsNullable {
		return bt.WithNullable()
	}
	return bt
}
