package bq

import (
	"encoding/json"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/emicklei/mtx/db"
	"github.com/emicklei/mtx/golang"
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
		if mt == mtx.Unknown {
			c.Set("maperror", each.AttributeType.Name)
		}
	}
	return tab
}

// ToBasicType returns a mapped basic Datatype
func ToBasicType(dt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(dt, registry.Class())

	if !dt.IsNullable {
		// expections
		if dt.Name == "DECIMAL" {
			return basic.Decimal.Set(golang.GoName, "*big.Rat")
		}
		if dt.Name == Date.Name {
			return basic.Date.Set(golang.GoName, "civil.Date")
		}
		if dt.Name == basic.JSON.Name {
			return basic.String
		}
		return *dt.BasicDatatype
	}
	var bt mtx.Datatype
	switch dt.Name {
	case String.Name:
		bt = basic.String.Set(golang.GoNullableTypeName, "bigquery.NullString")
	case "DECIMAL":
		bt = basic.Decimal.Set(golang.GoName, "*big.Rat").Nullable()
	case Bool.Name:
		bt = basic.Boolean.Set(golang.GoNullableTypeName, "bigquery.NullBool")
	case Timestamp.Name:
		bt = basic.Boolean.Set(golang.GoNullableTypeName, "bigquery.NullTimestamp")
	case Date.Name:
		bt = basic.Boolean.Set(golang.GoNullableTypeName, "bigquery.NullDate")
	case DateTime.Name:
		bt = basic.Boolean.Set(golang.GoNullableTypeName, "bigquery.NullDateTime")
	case DateTime.Name:
		bt = basic.Boolean.Set(golang.GoNullableTypeName, "bigquery.NullDateTime")
	case Bytes.Name:
		bt = basic.Bytes // empty bytes are considered null
	default:
		bt = basic.Unknown
	}
	return bt.Nullable()
}
