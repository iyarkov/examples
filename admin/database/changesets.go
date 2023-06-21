package database

import "github.com/iyarkov/foundation/schema"

var changeset = []schema.Change{
	{
		Version: "1.0.0",
		Commands: []string{
			`create table group_tbl (
				id int,
				created_at timestamp(3) without time zone not null,
				updated_at timestamp(3) without time zone not null,
				name varchar(255) constraint group_tbl_name_idx unique not null,
				primary key (id)
			)`,
		},
	},
}
