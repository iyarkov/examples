package database

import "github.com/iyarkov/foundation/schema"

var SchemaName = "public"

var expectedSchema = schema.Schema{
	Name: SchemaName,
	Tables: map[string]schema.Table{
		"group_tbl": {
			Columns: map[string]schema.Column{
				"id": {
					Type:         "int4",
					NotNull:      true,
					NumPrecision: 32,
				},
				"name": {
					Type:       "varchar",
					CharLength: 255,
					IsUnique:   true,
					NotNull:    true,
				},
				"created_at": {
					Type:    "timestamp",
					NotNull: true,
				},
				"updated_at": {
					Type:    "timestamp",
					NotNull: true,
				},
			},
			Indexes: map[string]schema.Index{
				"group_tbl_pkey": {
					Columns:  []string{"id"},
					IsUnique: true,
				},
				"group_tbl_name_idx": {
					Columns:  []string{"name"},
					IsUnique: true,
				},
			},
		},
	},
	Sequences: []string{
		"group_tbl_id_seq",
	},
}
