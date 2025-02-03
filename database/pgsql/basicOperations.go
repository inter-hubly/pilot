package pgsql

import "database/sql"

type SqlConn interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Row, error)
	QueryRows(query string, args ...any) (*sql.Rows, error)
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

func (c *connection) QueryRows(query string, args ...any) (*sql.Rows, error) {
	rows, err := c.pgsql.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
