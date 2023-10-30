# mtx

> **M**odels and **T**ransformations, E**x**pressed

A DSL to describe models targeting different technologies:

- postgres database
- gcp/bigquery database
- gcp/spanner database
- protocolbuffers
- golang struct
- csv
- entity–relationship

## example Postgres (pg) Database table

    db := pg.Database("example")
    tab := db.Table("persons")
    col := tab.Column("name").Datatype(pg.Text).Doc("what people call you")
    
    // or the short version
    col = tab.C("name",pg.Text,"what people call you")

Transforming this table to DDL creation statement in SQL

    sql := tab.ToSQL()

then sql:

        -- DROP TABLE persons
    CREATE TABLE persons (
        name text NOT NULL -- what people call you
    );

Transforming this table to an Entity representing a row.

    ent := tab.ToEntity()

The entity `ent` will have an Attribute `name` with Datatype `basic.String`

Transforming this entity to a Go struct

    str := golang.ToStruct(ent)

The struct `str` will have a Field `name` with Datatype `golang.String`

Transform this struct to Go source

    src := str.ToGo()

then src:

    // Persons :
    type Persons struct {
        Name string // what people call you
    }
## how to transform one model to another

A pg model (db.Table) first needs to be converted to an Entity model (basic.Entity).
The Entity model can be converted to a Go model (golang.Struct).
The Entity model can be modified to drive the conversion to the struct.

- rename
- change datatypes

Each model is using its own types and datatypes.
Transformations to and from a selection of model combinations.

