package database

import "github.com/iyarkov/foundation/sql"

var changeset = []sql.Change{
	{
		Version: "1.0.0",
		Commands: []string{
			`create table group_tbl (
				id serial,
				created_at timestamp(3) without time zone not null,
				updated_at timestamp(3) without time zone not null,
				name varchar(255) constraint group_tbl_name_idx unique not null,
				primary key (id)
			)`,
		},
	},
}
