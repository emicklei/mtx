# mtx

> **M**odels and **T**ransformations, E**x**pressed

A DSL to describe models targetting different technologies:

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

Transforming this table to SQL creation statement

    sql := tab.ToSQL()

Transforming this table to an Entity representing a row.

    ent := tab.ToEntity()

The entity `ent` will have an Attribute `name` with Datatype `basic.String`

Transforming this entity to a Go struct

    str := golang.ToStruct(ent)

The struct `str` will have a Field `name` with Datatype `golang.String`

Transform this struct to Go source

    src := str.ToGo()

## how to transform one model to another

A pg model (db.Table) first needs to be converted to an Entity model (basic.Entity).
The Entity model can be converted to a Go model (golang.Struct).
The Entity model can be modified to drive the convertion to the struct:

- rename
- change datatypes

Each model is using its own types and datatypes.
Transformations to and from a selection of model combinations.

