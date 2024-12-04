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