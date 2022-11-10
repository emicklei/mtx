package mtx

/**
string is non-nullable string
string? is a nullable string

string ->  mappings {
	golang: string
	bg: STRING,
	pg: pg.Text,
}
string? -> mappings {
	golang: *string
	bg: STRING,
	pg: pg.Text,
}

**/
