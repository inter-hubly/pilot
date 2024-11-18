package pgsql

import "database/sql"

type SqlConn interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Row, error)
}

func (c *connection) Exec(query string, args ...any) (sql.Result, error) {
	return c.pgsql.Exec(query, args...)
}

func (c *connection) Query(query string, args ...any) (*sql.Row, error) {
	row := c.pgsql.QueryRow(query, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}
