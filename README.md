# mtx

>   **M**odels and **T**ransformations, E**x**pressed

Models in different technologies:

- postgres database
- gcp/bigquery database
- gcp/spanner database
- protocolbuffers
- golang struct
- good old csv
- entity–relationship
 
## how to transform one model to another

A pg model (db.Table) first needs to be converted to an Entity model (basic.Entity).
The Entity model can be converted to a Go model (golang.Struct).
The Entity model can be modified to drive the convertion to the struct:
- rename
- change datatypes

Each model is using its own types and datatypes.
Transformations to and from a selection of model combinations.

## example Postgres (pg) Database table.

    db := pg.Database("example")
    tab := db.Table("persons")
    col := tab.Column("name").Datatype(pg.Text).Doc("what people call you")
    // or the short version
    col = tab.C("name",pg.Text,"what people call you")

Transforming this table to SQL creation statement

    sql := tab.ToSQL()

Transforming this table to an Entity representing a row.

    ent := tab.ToEntity()

The entity `ent` will have an Attribute `name` with Datatype `mtx.String`

Transforming this entity to a Go struct

    str := golang.ToStruct(ent)

The struct `str` will have a Field `name` with Datatype `golang.String`

Transform this struct to Go source

    src := str.ToGo()

### use case: Compose Proto messages from shared fields

### use case: SQL soruce to represent database table

### use case: Go source to represent row of database table

### use case: Compute diff

Save your model to JSON in a file with a version indicator.
If you change your model code later then you can compare it against the one stored.
Now you check for incompatible changes.

### use case: SQL alter table from diff

Compare a stored model with the current model of your database table.
Compute the diff and export the ALTER TABLE statement.

### use case: CSV into BigQuery

The csv package has a `ScanSheet` that builds a model from an existing CSV file for which the columns are typed based on the actual values from (enough) rows.
Transform this into a BigQuery model such that you can write code to insert rows into a BigQuery dataset table.

## test

    ./make.sh