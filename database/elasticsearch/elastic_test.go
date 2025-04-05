package elasticsearch

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/inter-hubly/pilot/server"
	"github.com/inter-hubly/pilot/testutils"
	"github.com/stretchr/testify/assert"
)

const (
	environmentDatabase = false
	elasticIndex        = "test.elastic"
)

func TestElastic(t *testing.T) {
	ctx := context.Background()

	var host string
	var close func(ctx context.Context) error
	var err error

	if environmentDatabase {
		host, close, err = testutils.ElasticSearch(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if close != nil {
			defer close(ctx)
		}
	} else {
		os.Setenv("ENVIRONMENT", "test")
		server.MockStartEnv(ctx, "../../")
		host = server.GetElasticSearch().Host
	}

	type elastic struct {
		conn *connection
	}
	NewConn(WithUrl([]string{host}))
	repo := elastic{
		conn: GetConnection(),
	}

	// t.Run("need be saved", func(t *testing.T) {
	// 	docT := newDocumentTest()
	// 	resp := &Response{}
	//
	// 	resp, err = repo.conn.Create(ctx, elasticIndex, docT)
	// 	assert.Empty(t, err)
	// 	assert.NotNil(t, resp.ID)
	// 	assert.NotNil(t, resp.Result)
	//
	// 	response, err2 := repo.conn.FindById(ctx, elasticIndex, resp.ID)
	// 	assert.Nil(t, err2)
	// 	assert.NotNil(t, response)
	// 	assert.Equal(t, resp.ID, response.ID)
	//
	// 	query := map[string]interface{}{
	// 		"query": map[string]interface{}{
	// 			"bool": map[string]interface{}{
	// 				"must": []map[string]interface{}{
	// 					{
	// 						"ids": map[string]interface{}{
	// 							"values": []string{response.ID},
	// 						},
	// 					},
	// 					{
	// 						"match": map[string]interface{}{
	// 							"test": "myTest",
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	}
	//
	// 	update, err3 := repo.conn.Update(ctx, elasticIndex, query)
	// 	assert.Nil(t, err3)
	// 	assert.NotNil(t, update)
	// 	assert.Equal(t, 1, update.Updated)
	// })

	t.Run("need find all", func(t *testing.T) {

		for i := 0; i < 3; i++ {
			docT := newDocumentTest()
			docT.Id = fmt.Sprintf("%s%d", docT.Id, i)
			_, err := repo.conn.Create(ctx, elasticIndex, docT)
			assert.NoError(t, err)
		}

		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}

		response, err2 := repo.conn.FindAll(ctx, elasticIndex, query)
		assert.NoError(t, err2)
		assert.NotNil(t, response.Hits)
	})
}

type documentTest struct {
	Id   string `json:"id"`
	Test string `json:"test"`
}

func newDocumentTest() *documentTest {
	return &documentTest{
		Id:   "123456789",
		Test: "myTest",
	}
}
