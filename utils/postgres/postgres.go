package postgres

import "database/sql"

func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}
