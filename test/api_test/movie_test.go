package api_test

import (
	"interface_project/test"
	"testing"
)

func TestMovieAPI(t *testing.T) {
	testClient := test.GetTestClientAndContext(t)
	defer testClient.CallCancelAndClose()
	// addMovieTest(*testClient)
}

func addMovieEndpoint(tc test.TestClient) {
	
}
