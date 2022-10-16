// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// RanobesColumns holds the columns for the "ranobes" table.
	RanobesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "name", Type: field.TypeString},
		{Name: "eng_name", Type: field.TypeString},
		{Name: "rus_name", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString, Size: 2147483647},
		{Name: "release_date", Type: field.TypeInt},
	}
	// RanobesTable holds the schema information for the "ranobes" table.
	RanobesTable = &schema.Table{
		Name:       "ranobes",
		Columns:    RanobesColumns,
		PrimaryKey: []*schema.Column{RanobesColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "priority", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_users_todos",
				Columns:    []*schema.Column{TodosColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// RanobeCategoriesColumns holds the columns for the "ranobe_categories" table.
	RanobeCategoriesColumns = []*schema.Column{
		{Name: "ranobe_id", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeString},
	}
	// RanobeCategoriesTable holds the schema information for the "ranobe_categories" table.
	RanobeCategoriesTable = &schema.Table{
		Name:       "ranobe_categories",
		Columns:    RanobeCategoriesColumns,
		PrimaryKey: []*schema.Column{RanobeCategoriesColumns[0], RanobeCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ranobe_categories_ranobe_id",
				Columns:    []*schema.Column{RanobeCategoriesColumns[0]},
				RefColumns: []*schema.Column{RanobesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ranobe_categories_category_id",
				Columns:    []*schema.Column{RanobeCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		RanobesTable,
		TodosTable,
		UsersTable,
		RanobeCategoriesTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = UsersTable
	RanobeCategoriesTable.ForeignKeys[0].RefTable = RanobesTable
	RanobeCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
}
