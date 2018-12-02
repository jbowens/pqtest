package pgtest

import "testing"

func TestOpen(t *testing.T) {
	db := Open(t)
	err := db.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenSchema(t *testing.T) {
	db := Open(t, SchemaFile("example.sql"))
	err := db.Close()
	if err != nil {
		t.Fatal(err)
	}
}
