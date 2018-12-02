/*
Package pqtest provides utilities for using
PostgreSQL databases in tests.

Open

Open creates a new PostgreSQL database initialized with
a schema and returns an opened *sql.DB connected to the
new database.

	func TestSomething(t *testing.T) {
		db := pqtest.Open(t, pqtest.SchemaFile("example.sql"))
		defer db.Close()

		// ...
	}



Removing databases

Each call to pqtest.Open queries the PostgreSQL catalog
for databases created by pqtest in the past. If there
are old databases, Open deletes up to five of them on
each invocation.

*/
package pqtest
