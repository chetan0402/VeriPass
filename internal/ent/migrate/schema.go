// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 26},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "hostel", Type: field.TypeString},
		{Name: "can_add_pass", Type: field.TypeBool},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
	}
	// PassesColumns holds the columns for the "passes" table.
	PassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"unspecified", "class", "market", "home", "event"}, Default: "unspecified"},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
	}
	// PassesTable holds the schema information for the "passes" table.
	PassesTable = &schema.Table{
		Name:       "passes",
		Columns:    PassesColumns,
		PrimaryKey: []*schema.Column{PassesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "room", Type: field.TypeString},
		{Name: "hostel", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		PassesTable,
		UsersTable,
	}
)

func init() {
}
