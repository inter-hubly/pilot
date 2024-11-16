package pgsql

import (
	"testing"

	"github.com/inter-hubly/pilot/testutils"
	"github.com/stretchr/testify/assert"
)

func TestPgsql(t *testing.T) {
	var err error
	container := testutils.NewPgsqlTestContainer()
	if err = container.Pgsql(); err != nil {
		t.Fatal(err)
	}
	NewConnection(WithUrl(container.GetHost()))

	conn := GetConnection()
	t.Run("need insert test", func(t *testing.T) {
		one, insertError := conn.pgsql.Exec(`
	    CREATE TABLE IF NOT EXISTS test (
	        id SERIAL PRIMARY KEY,
	        name TEXT NOT NULL
	    );
	`)
		assert.NoError(t, err)

		v, err := conn.pgsql.Exec("INSERT INTO test (name) VALUES ($1)", "test")
		assert.NoError(t, err)

		assert.NoError(t, insertError)
		assert.NotNil(t, one)
		assert.NotNil(t, v)

		var name string
		err = conn.pgsql.QueryRow("SELECT name FROM test WHERE name = $1", "test").Scan(&name)
		assert.NoError(t, err)
		assert.Equal(t, "test", name)

	})

	for _, v := range []struct {
		host      string
		port      string
		database  string
		username  string
		password  string
		urlString string
	}{
		{
			host: "localhost", port: "1233", database: "testdb", username: "username", password: "password",
			urlString: "postgres://username:password@localhost:1233/testdb?sslmode=disable",
		},
	} {
		t.Run("extract url values", func(t *testing.T) {
			host, port, database, username, password, err := extractDBDetails(v.urlString)

			assert.NoError(t, err)
			assert.Equal(t, v.host, host)
			assert.Equal(t, v.port, port)
			assert.Equal(t, v.database, database)
			assert.Equal(t, v.username, username)
			assert.Equal(t, v.password, password)

		})
	}

}
