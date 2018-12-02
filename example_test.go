package pqtest_test

import (
	"testing"

	"github.com/jbowens/pqtest"
)

func ExampleSchemaFile(t *testing.T) {
	db := pqtest.Open(t, pqtest.SchemaFile("example.sql"))
	defer db.Close()
	// ...
}
