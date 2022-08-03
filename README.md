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
 
Each model is using its own types and datatypes.
Transformations to and from a selection of model combinations.


With relational databases, "Table" and "Column" are the building blocks that exist in DDL 
and it's best to just use those terms and avoid "field" which isn't used, nor clearly defined.


Example Postgres (pg) Database table.

    db := pg.Database("example")
    tab := db.Table("persons")
    col := tab.Column("name").Datatype(pg.TEXT)

Transforming this table to SQL creation statement

    sql := tab.SQL()

Transforming this table to an Entity representing a row.

    ent := tab.ToEntity()

The entity `ent` will have an Attribute "name" with Datatype mtx.STRING

Transforming this entity to a Go struct

    str := golang.ToStruct(ent)

The struct `str` will have a Field "name" with Datatype golang.STRING

Transform this struct to Go source

    src := str.Go()

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