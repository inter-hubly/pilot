package elasticsearch

import (
	"context"
	"log"
	"testing"

	"github.com/inter-hubly/pilot/testutils"
	"github.com/stretchr/testify/assert"
)

func TestElastic(t *testing.T) {

	const elasticIndex = "test.elastic"
	ctx := context.Background()
	host, close, err := testutils.ElasticSearch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if close != nil {
		defer close(ctx)
	}

	type elastic struct {
		conn *connection
	}
	NewConn(WithUrl([]string{host}))
	repo := elastic{
		conn: GetConnection(),
	}

	for _, v := range []struct {
		testName string
		isErr    bool
	}{
		{
			testName: "need save value",
			isErr:    false,
		},
	} {
		t.Run(v.testName, func(t *testing.T) {
			docT := newDocumentTest()
			resp := &Response{}

			resp, err = repo.conn.Create(ctx, elasticIndex, docT)
			if v.isErr {
				log.Print(err)
				assert.NotEmpty(t, err)
				assert.Nil(t, resp)
				return
			}
			assert.Empty(t, err)
			assert.NotNil(t, resp.ID)
			assert.NotNil(t, resp.Result)

			response, err2 := repo.conn.FindById(ctx, elasticIndex, resp.ID)
			assert.Nil(t, err2)
			assert.NotNil(t, response)
			assert.Equal(t, resp.ID, response.ID)
		})
	}
}

type documentTest struct {
	Id   string `json:"id"`
	Test string `json:"test"`
}

func newDocumentTest() *documentTest {
	return &documentTest{
		Test: "myTest",
	}
}
