package database

import "github.com/iyarkov/foundation/sql"

var SchemaName = "public"

var expectedSchema = sql.Schema{
	Name: SchemaName,
	Tables: map[string]sql.Table{
		"group_tbl": {
			Columns: map[string]sql.Column{
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
			Indexes: map[string]sql.Index{
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
