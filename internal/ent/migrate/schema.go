// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
	}
)

func init() {
}
