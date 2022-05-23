package test

import (
	"context"
	"interface_project/api"
	"interface_project/database"
	"interface_project/ent"
	"interface_project/ent/enttest"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type TestClient struct {
	context.Context
	ent.Client
	testing.T
	context.CancelFunc
	api.API
	*httptest.ResponseRecorder
}

func (testClient TestClient) CallCancelAndClose() {
	testClient.CancelFunc()
	testClient.Client.Close()
}

func GetTestClientAndContext(t *testing.T) *TestClient {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx, cancel := database.GetDatabaseContext()
	if err := client.Schema.WriteTo(*ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %+v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema changes: %+v", err)
	}
	router := api.RunAPI(ctx, client)
	w := httptest.NewRecorder()
	return &TestClient{Context: *ctx, Client: *client, T: *t, CancelFunc: *cancel, API: *router, ResponseRecorder: w}
}
